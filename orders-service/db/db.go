package db

import (
	"github.com/jinzhu/gorm"
	"github.com/seb7887/go-microservices/config"
	"github.com/seb7887/go-microservices/models"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var DB *gorm.DB

func InitDatabase() {
	database, err := gorm.Open(config.GetConfig().DBType, config.GetConfig().DBConnection)
	if err != nil {
		log.Fatal(err)
	}

	// Set up connection pool
	database.DB().SetMaxIdleConns(20)
	database.DB().SetMaxOpenConns(200)
	DB = database
}

func AutoMigrate() error {
	err := DB.AutoMigrate(&models.Order{}).Error
	if err != nil {
		log.Fatal(err)
	}
	return nil
}