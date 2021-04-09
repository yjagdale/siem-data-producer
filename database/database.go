package database

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var Connection *gorm.DB

func ConnectToDb() {
	newLogger := logger.New(
		log.New(),
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // Disable color
		},
	)
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatalln("Error while connecting to db. ", err.Error())
	}
	Connection = db
}

func ValidateDBConnection() bool {
	err := Connection.Exec("SELECT 1").Error
	return err == nil
}
