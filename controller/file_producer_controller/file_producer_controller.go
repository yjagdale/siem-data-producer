package file_producer_controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/yjagdale/siem-data-producer/models/file_producer_model"
	"github.com/yjagdale/siem-data-producer/services/file_producer_service"
	"github.com/yjagdale/siem-data-producer/utils/response"
	"io"
	"net/http"
	"os"
)

func ProduceFile(c *gin.Context) {
	var fileEntity file_producer_model.FileProducer
	isValid := c.ShouldBindJSON(&fileEntity)

	if isValid == io.EOF {
		log.Errorln(isValid)
		resp := response.NewBadRequest(gin.H{"Message": "Empty Body? May be!"})
		c.JSON(resp.Status, resp)
		return
	}

	if isValid != nil {
		log.Errorln(isValid)
		resp := response.NewBadRequest(gin.H{"Message": "Invalid Body"})
		c.JSON(resp.Status, resp)
		return
	}

	log.Infoln("Producing logs on destination", fileEntity.DestinationIP, "over port", fileEntity.DestinationPort, "From file", fileEntity.Path)
	stats, err := os.Stat(fileEntity.Path)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	} else {
		if stats.Size() > 1594682 {
			go file_producer_service.PublishFile(fileEntity)
			c.JSON(http.StatusAccepted, gin.H{"Message": "Large file provided. Execution will be done in background"})
		} else {
			resp := file_producer_service.PublishFile(fileEntity)
			c.JSON(resp.Status, resp.Message)
			return
		}
	}
}
