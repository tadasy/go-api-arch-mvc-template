package configs

import (
	"go-api-arch-mvc-template/pkg/logger"
	"os"
	"strconv"

	"go.uber.org/zap"
)

func GetEnvDefault(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

type ConfigList struct {
	Env                 string
	DBHost              string
	DBPort              int
	DBDriver            string
	DBName              string
	DBUser              string
	DBPassword          string
	APICorsAllowOrigins []string
}

func (c *ConfigList) IsDevelopment() bool {
	return c.Env == "development"
}

var Config ConfigList

func LoadEnv() error {
	DBPort, err := strconv.Atoi(GetEnvDefault("MYSQL_PORT", "3306"))
	if err != nil {
		return err
	}
	Config = ConfigList{
		Env:                 GetEnvDefault("APP_ENV", "development"),
		DBDriver:            GetEnvDefault("DB_DRIVER", "mysql"),
		DBHost:              GetEnvDefault("DB_HOST", "0.0.0.0"),
		DBPort:              DBPort,
		DBUser:              GetEnvDefault("DB_USER", "app"),
		DBPassword:          GetEnvDefault("DB_PASSWORD", "password"),
		DBName:              GetEnvDefault("DB_NAME", "api_database"),
		APICorsAllowOrigins: []string{"http://0.0.0.0:8001"},
	}
	return nil
}

func init() {
	if err := LoadEnv(); err != nil {
		logger.Error("Failed to load env: ", zap.Error(err))
		panic(err)
	}
}
