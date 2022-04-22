package main

import (
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/nightlord189/so5hw/internal/handler"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nightlord189/so5hw/docs"
	"github.com/nightlord189/so5hw/internal/config"
	"github.com/nightlord189/so5hw/internal/db"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"time"
)

// @title so5hw
// @description backend
// @schemes http https
// @version 1.0
// @BasePath /
func main() {
	fmt.Println("start")
	cfg, err := config.Load("configs/config.json")
	if err != nil {
		panic(fmt.Sprintf("error initializing config file: %v", err))
	}

	dbInstance, err := db.InitDb(cfg)
	if err != nil {
		panic(fmt.Sprintf("err init db: %v", err))
	}
	handlerInst := handler.NewHandler(cfg, dbInstance)
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Access-Control-Allow-Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Range"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))

	router.Use(static.Serve("/", static.LocalFile("./web/build", true)))

	//for React-router (non-SPA)
	router.NoRoute(func(c *gin.Context) {
		if !strings.Contains(c.Request.RequestURI, "/api/") {
			c.File("./web/build/index.html")
		}
	})

	router.GET("/healthz", func(c *gin.Context) {
		c.String(200, "success")
	})
	router.GET("/swagger/*any", func(context *gin.Context) {
		docs.SwaggerInfo.BasePath = cfg.SwaggerBasePath
		ginSwagger.WrapHandler(swaggerFiles.Handler)(context)
	})

	api := router.Group("/api")

	service := api.Group("/service")
	service.POST("/reset", handlerInst.ResetDB)
	service.POST("/fill", handlerInst.FillDB)

	authMiddleware := handler.CheckAuthMiddleware(cfg.AuthAccessSecret)

	auth := api.Group("/auth")
	auth.POST("", handlerInst.Auth)

	customer := api.Group("/customer", authMiddleware)
	customer.GET("/:id", handlerInst.GetCustomer)

	merchandiser := api.Group("/merchandiser", authMiddleware)
	merchandiser.GET("/:id", handlerInst.GetMerchandiser)

	product := api.Group("/product", authMiddleware)
	product.GET("", handlerInst.GetProducts)
	product.GET("/category", handlerInst.GetCategories)
	product.POST("/:id", handlerInst.CreateProduct)
	product.DELETE("/:id", handlerInst.DeleteProduct)

	sale := api.Group("/sale", authMiddleware)
	sale.POST("", handlerInst.Sale)

	err = dbInstance.TruncateAllTables()
	if err != nil {
		fmt.Printf("err truncate tables: %v\n", err)
	}
	err = dbInstance.FillData()
	if err != nil {
		fmt.Printf("err fill data: %v\n", err)
	}

	err = router.Run(fmt.Sprintf(":%d", cfg.HTTPPort))
	if err != nil {
		panic(fmt.Sprintf("err run router: %v", err))
	}
}
