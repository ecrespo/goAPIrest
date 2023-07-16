package controllers

import (
	"github.com/ecrespo/goAPIrest/api/database"
	"github.com/ecrespo/goAPIrest/api/models"
	"github.com/ecrespo/goAPIrest/api/repositories"
	"github.com/ecrespo/goAPIrest/api/utils/logs"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

type Server struct {
	DB             *gorm.DB
	Router         *mux.Router
	UserRepository repositories.UserRepository
	PostRepository repositories.PostRepository
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
		log.Fatalf("Error while initializing database: %v\n", err)
	}

	server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{}) //database migration

	server.Router = mux.NewRouter()

	server.UserRepository = repository.NewUserRepository(server.DB)
	server.PostRepository = repositories.NewPostRepository(server.DB)

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	logger := logs.GetLogger()
	logger.Info().Msgf("Listening to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
