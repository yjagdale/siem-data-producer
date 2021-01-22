package app

import (
	"github.com/gin-gonic/gin"
	"github.com/yjagdale/siem-data-producer/controller/files_controller"
	"net/http"
)

func MapUrls() {
	router = gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})
	router.Static("/ui", "./static")

	/* file controller */

	router.POST("/upload", files_controller.Upload)

}
