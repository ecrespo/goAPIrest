package controllers

import (
	"github.com/ecrespo/goAPIrest/api/database"
	"github.com/ecrespo/goAPIrest/api/models"
	"github.com/ecrespo/goAPIrest/api/utils/logs"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error
	db := &database.Database{ // Aquí se cambió la estructura de Database
		Driver:   Dbdriver,
		User:     DbUser,
		Password: DbPassword,
		Port:     DbPort,
		Host:     DbHost,
		Name:     DbName,
	}

	server.DB, err = db.Initialize()
	if err != nil {
		// handle error
	}

	server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{}) //database migration

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	logger := logs.GetLogger()
	logger.Info().Msgf("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
