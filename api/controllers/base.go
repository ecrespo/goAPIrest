package controllers

import (
	"fmt"
	"github.com/ecrespo/goAPIrest/api/utils/logs"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql database driver
	//_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
	"github.com/ecrespo/goAPIrest/api/models"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error
	logger := logs.GetLogger()
	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			//fmt.Printf("Cannot connect to %s database", Dbdriver)
			logger.Info().Msgf("Cannot connect to %s database", Dbdriver)
			logger.Fatal().Msgf("This is the error:", err)
			//log.Fatal("This is the error:", err)
		} else {
			logger.Info().Msgf("We are connected to the %s database", Dbdriver)
			//fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}
	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			//fmt.Printf("Cannot connect to %s database", Dbdriver)
			logger.Info().Msgf("Cannot connect to %s database", Dbdriver)
			//log.Fatal("This is the error:", err)
			logger.Fatal().Msgf("This is the error:", err)
		} else {
			//fmt.Printf("We are connected to the %s database", Dbdriver)
			logger.Info().Msgf("We are connected to the %s database", Dbdriver)
		}
	}

	server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{}) //database migration

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	logger := logs.GetLogger()
	//fmt.Println("Listening to port 8080")
	logger.Info().Msgf("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))

}
