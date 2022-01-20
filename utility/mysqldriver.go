package utility

import (
	"os"
	_config "tixid/config"
	_entity "tixid/entity"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *_config.AppConfig) *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	connectionString := os.Getenv("CONNECTION_DB")
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Info("failed to connect database", err)
		panic(err)
	}
	InitMigrate(db)
	return db
}

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(_entity.User{})
}
