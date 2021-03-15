package producer_controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/yjagdale/siem-data-producer/models/model_logs"
	"github.com/yjagdale/siem-data-producer/services/producer_service"
	"github.com/yjagdale/siem-data-producer/utils/error_response"
	"io"
)

func Produce(c *gin.Context) {

	var producerEntity model_logs.ProducerEntity

	isValid := c.ShouldBindJSON(&producerEntity)

	if isValid == io.EOF {
		log.Errorln(isValid)
		response := error_response.NewBadRequest("Empty Body? May be!")
		c.JSON(response.Status, response)
		return
	}

	if isValid != nil {
		log.Errorln(isValid)
		response := error_response.NewBadRequest("Invalid Body")
		c.JSON(response.Status, response)
		return
	}

	log.Infoln("Producing logs on destination", producerEntity.DestinationIP, "over port", producerEntity.DestinationPort)
	err := producer_service.ProduceSingle(producerEntity)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(202, gin.H{"Message": "Producer running in background", "status": 202})
}
