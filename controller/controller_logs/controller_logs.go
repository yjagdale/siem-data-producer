package controller_logs

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seca-data-producer/services/services_logs"
)

func GetLogFiles(c *gin.Context) {
	resp, err := services_logs.GetAvailableLogFiles()

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, resp)
	return
}

func GetLogFile(c *gin.Context) {
	logsFileName := c.Param("file_name")
	resp, err := services_logs.GetAvailableLogFiles(logsFileName)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, resp)
	return
}
