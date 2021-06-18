package app

import (
	"github.com/gin-gonic/gin"
	"github.com/yjagdale/siem-data-producer/controller/file_producer_controller"
	"github.com/yjagdale/siem-data-producer/controller/files_controller"
	"github.com/yjagdale/siem-data-producer/controller/override_controller"
	"github.com/yjagdale/siem-data-producer/controller/producer_controller"
	"net/http"
)

func MapUrls() {
	router = gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})
	/* file controller */

	router.POST("/upload", files_controller.Upload)

	router.Static("/ui", "./static")

	router.POST("/producer/produce", producer_controller.Produce)
	router.POST("/producer/produce/async", producer_controller.ProduceAsync)
	router.POST("/producer/produce/execute", producer_controller.ProduceContinues)
	router.POST("/producer/produce/execute/stop/:executionId", producer_controller.StopExecutor)

	router.GET("/producer/produce/executions", producer_controller.GetExecutions)

	router.POST("/producer/produce/test", producer_controller.ProduceTest)

	router.POST("/producer/produce/file", file_producer_controller.ProduceFile)
	router.POST("/producer/produce/file/async", producer_controller.ProduceAsync)
	router.POST("/producer/produce/file/test", producer_controller.ProduceTest)

	router.POST("/producer/overrides", override_controller.AddOverride)

}
