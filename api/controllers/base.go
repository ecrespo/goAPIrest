package controllers

import (
	"github.com/ecrespo/goAPIrest/api/database"
	"github.com/ecrespo/goAPIrest/api/utils/logs"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	logger := logs.GetLogger()

	server.DB, err = database.NewDatabase(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName)
	if err != nil {
		logger.Fatal().Msgf("Failed to connect to database: %v", err)
	}

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	logger := logs.GetLogger()
	logger.Info().Msgf("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
