// File: server.go
package api

import (
	"github.com/ecrespo/goAPIrest/api/controllers"
	"github.com/ecrespo/goAPIrest/api/seed"
	"github.com/ecrespo/goAPIrest/api/utils/logs"
	"github.com/joho/godotenv"
	"os"
)

// Instance of the server
var server = controllers.Server{}

func Run() {
	if err := loadEnvVariables(); err != nil {
		logs.GetLogger().Fatal().Msgf("Failed to load environment variables: %v", err)
		return
	}

	initializeServer()
	startApplication()
}

func loadEnvVariables() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	logs.GetLogger().Info().Msg("Environment variables loaded successfully")
	return nil
}

func initializeServer() {
	dbEnv := getDBEnvironmentVariables()
	server.Initialize(dbEnv.driver, dbEnv.user, dbEnv.password, dbEnv.port, dbEnv.host, dbEnv.name)
	seed.Load(server.DB)
}

func startApplication() {
	logs.GetLogger().Info().Msg("Starting the application...")
	server.Run(":8080")
}

type dbEnvironment struct {
	driver   string
	user     string
	password string
	port     string
	host     string
	name     string
}

func getDBEnvironmentVariables() dbEnvironment {
	return dbEnvironment{
		driver:   os.Getenv("DB_DRIVER"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		port:     os.Getenv("DB_PORT"),
		host:     os.Getenv("DB_HOST"),
		name:     os.Getenv("DB_NAME"),
	}
}
