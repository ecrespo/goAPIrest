package controllers

import (
	"fmt"
	"github.com/ecrespo/goAPIrest/api/utils/logs"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	"github.com/ecrespo/goAPIrest/api/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type Database struct {
	Driver, User, Password, Port, Host, Name string
}

func (db *Database) Initialize() (*gorm.DB, error) {
	var connectionStr string
	logger := logs.GetLogger()

	switch db.Driver {
	case "mysql":
		connectionStr = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", db.User, db.Password, db.Host, db.Port, db.Name)
	case "postgres":
		connectionStr = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", db.Host, db.Port, db.User, db.Name, db.Password)
	default:
		return nil, fmt.Errorf("Unsupported driver: %s", db.Driver)
	}

	conn, err := gorm.Open(db.Driver, connectionStr)
	if err != nil {
		logger.Info().Msgf("Cannot connect to %s database", db.Driver)
		logger.Fatal().Msgf("This is the error:", err)
	}
	logger.Info().Msgf("We are connected to the %s database", db.Driver)

	return conn, err
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error
	db := &Database{
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
