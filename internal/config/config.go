package config

import (
	"encoding/json"
	"os"

	"github.com/kelseyhightower/envconfig"
)

// Config - configuration struct
type Config struct {
	HTTPPort         int    `envconfig:"HTTP_PORT"`
	GinMode          string // debug or release
	TokenExpTime     int
	AuthAccessSecret string `envconfig:"AUTH_ACCESS_SECRET"`
	SwaggerBasePath  string `envconfig:"SWAGGER_BASE_PATH"`
	DB               struct {
		Host     string `envconfig:"DB_HOST"`
		Port     int    `envconfig:"DB_PORT"`
		Name     string `envconfig:"DB_NAME"`
		User     string `envconfig:"DB_USER"`
		Password string `envconfig:"DB_PASSWORD"`
		Migrate  bool
		Log      bool `envconfig:"DB_LOG"`
	}
}

// Load - load config from path
func Load(path string) (*Config, error) {
	config := Config{}
	file, err := os.Open(path)
	if err != nil {
		panic("load config error: " + err.Error())
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return &config, err
	}
	err = envconfig.Process("", &config)
	return &config, err
}
