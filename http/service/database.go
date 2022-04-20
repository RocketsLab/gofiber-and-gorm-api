package service

import (
	"github.com/RocketsLab/gofiber-and-gorm-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	DbConnection *gorm.DB
)

func InitDatabase() {
	//logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	var err error
	DbConnection, err = gorm.Open(sqlite.Open("database/database.db"), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic(err.Error())
	}

	err = DbConnection.AutoMigrate(&models.User{})
	if err != nil {
		panic(err.Error())
	}
}
