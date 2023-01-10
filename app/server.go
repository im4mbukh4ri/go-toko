package app

import (
	"flag"
	"log"
	"os"

	"github.com/im4mbukh4ri/go-toko/app/controllers"
	"github.com/joho/godotenv"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Run() {
	var server = controllers.Server{}
	var appConfig = controllers.AppConfig{}
	var dbConfig = controllers.DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error on loading .env file")
	}

	appConfig.AppName = getEnv("APP_NAME", "Go Toko")
	appConfig.AppEnv = getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT", "9000")

	dbConfig.DBHost = getEnv("DB_HOST", "localhost")
	dbConfig.DBUser = getEnv("DB_USER", "root")
	dbConfig.DBPassword = getEnv("DB_PASSWORD", "secret")
	dbConfig.DBName = getEnv("DB_NAME", "db-name")
	dbConfig.DBPort = getEnv("DB_PORT", "5432")
	dbConfig.DBTimezone = getEnv("APP_TIMEZONE", "Asia/Jakarta")
	dbConfig.DBDriver = getEnv("DB_DRIVER", "postgres")
	flag.Parse()
	arg := flag.Arg(0)
	if arg != "" {
		server.InitCommands(appConfig, dbConfig)
	} else {
		server.Initialize(appConfig, dbConfig)
		server.Run(":" + appConfig.AppPort)
	}

}
