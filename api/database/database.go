package database

import (
	"fmt"
	"github.com/ecrespo/goAPIrest/api/utils/logs"
	"github.com/jinzhu/gorm"
)

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
