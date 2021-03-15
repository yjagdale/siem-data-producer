package app

/* import (
	"github.com/yjagdale/siem-data-producer/database"
	"github.com/yjagdale/siem-data-producer/models/model_logs"
	"log"
)

func MigrateDB() {
	database.ConnectToDb()

	if database.Connection != nil {
		err := database.Connection.AutoMigrate(&model_logs.AvailableLogs{}, &model_logs.Logs{})
		if err != nil {
			log.Fatalln("Error while migrating database. ", err.Error())
		}
	}
}
*/
