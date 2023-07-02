package database

import (
	"fmt"
	"github.com/ecrespo/goAPIrest/api/models"
	"github.com/ecrespo/goAPIrest/api/utils/logs"
	"github.com/jinzhu/gorm"
)

func NewDatabase(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*gorm.DB, error) {
	var DB *gorm.DB
	var err error
	logger := logs.GetLogger()

	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			logger.Info().Msgf("Cannot connect to %s database", Dbdriver)
			logger.Fatal().Msgf("This is the error:", err)
		} else {
			logger.Info().Msgf("We are connected to the %s database", Dbdriver)
		}
	}

	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			logger.Info().Msgf("Cannot connect to %s database", Dbdriver)
			logger.Fatal().Msgf("This is the error:", err)
		} else {
			logger.Info().Msgf("We are connected to the %s database", Dbdriver)
		}
	}

	DB.Debug().AutoMigrate(&models.User{}, &models.Post{}) //database migration

	return DB, nil
}
