package app

import (
	"github.com/gin-gonic/gin"
	"github.com/yjagdale/siem-data-producer/utils/logger"
	"log"
)

var router *gin.Engine

func StartApp() {
	logger.InitLogger()
	MigrateDB()
	MapUrls()
	startApplication()
}

func startApplication() {
	err := router.Run(":8080")
	if err != nil {
		log.Fatalln("failed to start ", err.Error())
	}
}
