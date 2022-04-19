package handler

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/nightlord189/so5hw/internal/config"
	"github.com/nightlord189/so5hw/internal/db"
	"github.com/ory/dockertest/v3"
	golog "log"
	"os"
	"strconv"
	"strings"
	"testing"
)

var configInstance *config.Config
var dbInstance *db.Manager
var handlerInstance *Handler

func TestMain(m *testing.M) {
	teardown := setupTests()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setupTests() func() {
	configInstance, _ = config.Load("../../configs/config.json")
	teardown := startDb(configInstance)
	gin.SetMode(gin.TestMode)
	dbInst, err := db.InitDb(configInstance)
	if err != nil {
		panic(fmt.Sprintf("error init db: %v", err))
	}
	dbInstance = dbInst
	handlerInstance = NewHandler(configInstance, dbInstance)
	return teardown
}

//startDb - launch new test db instance in docker container
func startDb(conf *config.Config) func() {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		golog.Fatalf("Could not connect to docker: %v", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("postgres", "12-alpine", []string{
		"POSTGRES_DB=" + conf.DB.Name,
		"POSTGRES_PASSWORD=" + conf.DB.Password,
		"POSTGRES_USER=" + conf.DB.User,
		"listen_addresses = '*'",
	})
	if err != nil {
		golog.Fatalf("Could not start resource: %s", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")

	port := strings.Split(hostAndPort, ":")[1]
	portInt, _ := strconv.Atoi(port)
	//on gitlab CI/CD pipeline host should be "docker" from env variable
	//on local it should be "localhost"
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error
		databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", conf.DB.User, conf.DB.Password, dbHost, port, conf.DB.Name)
		golog.Println("Connecting to database on url: ", databaseUrl)
		dbInstance2, err := sql.Open("postgres", databaseUrl)
		if err != nil {
			return err
		}
		return dbInstance2.Ping()
	}); err != nil {
		golog.Fatalf("Could not connect to database: %s", err)
	}

	//change config db values to new created docker
	conf.DB.Port = portInt
	conf.DB.Host = dbHost

	fmt.Println("success start docker db")

	return func() {
		if err := pool.Purge(resource); err != nil {
			golog.Fatalf("Could not purge resource: %s", err)
			return
		}
		fmt.Println("docker purged")
	}
}
