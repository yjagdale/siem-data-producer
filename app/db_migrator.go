package app

import (
	"github.com/yjagdale/siem-data-producer/database"
	"github.com/yjagdale/siem-data-producer/models/override_model"
	"log"
)

func MigrateDB() {
	database.ConnectToDb()

	if database.Connection != nil {
		err := database.Connection.AutoMigrate(&override_model.OverrideConfig{})
		if err != nil {
			log.Fatalln("Error while migrating database. ", err.Error())
		}
	}
}
