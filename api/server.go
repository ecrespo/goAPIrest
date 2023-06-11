package api

import (
	"github.com/ecrespo/goAPIrest/api/controllers"
	"github.com/ecrespo/goAPIrest/api/seed"
	"github.com/ecrespo/goAPIrest/api/utils/logs"
	"github.com/joho/godotenv"
	//"log"
	"os"
	//"log"
)

var server = controllers.Server{}

func Run() {

	var err error
	err = godotenv.Load()
	logger := logs.GetLogger()
	if err != nil {
		logger.Fatal().Msgf("Error getting env, not comming through %v", err)
		//log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		logger.Info().Msg("We are getting the env values")
		//fmt.Println("We are getting the env values")
	}
	logger.Info().Msgf("DB_DRIVER: %s", os.Getenv("DB_DRIVER"))
	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	seed.Load(server.DB)
	logger.Info().Msg("About to start the application...")
	server.Run(":8080")

}
