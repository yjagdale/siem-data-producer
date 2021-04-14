package producer_controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/yjagdale/siem-data-producer/config/constant"
	"github.com/yjagdale/siem-data-producer/models/producer_model"
	"github.com/yjagdale/siem-data-producer/services/producer_service"
	"github.com/yjagdale/siem-data-producer/utils/response"
	"io"
	"net/http"
)

func Produce(c *gin.Context) {

	var producerEntity producer_model.ProducerEntity

	isValid := c.ShouldBindJSON(&producerEntity)

	if isValid == io.EOF {
		log.Errorln(isValid)
		response := response.NewBadRequest(gin.H{"Message": "Empty Body? May be!"})
		c.JSON(response.Status, response)
		return
	}

	if isValid != nil {
		log.Errorln(isValid)
		response := response.NewBadRequest(gin.H{"Message": "Invalid Body"})
		c.JSON(response.Status, response)
		return
	}

	log.Infoln("Producing logs on destination", producerEntity.DestinationIP, "over port", producerEntity.DestinationPort)
	response := producer_service.Produce(producerEntity, constant.ExecutionModeProduce)
	c.JSON(response.Status, response)
}

func ProduceAsync(c *gin.Context) {

	var producerEntity producer_model.ProducerEntity

	isValid := c.ShouldBindJSON(&producerEntity)

	if isValid == io.EOF {
		log.Errorln(isValid)
		restResponse := response.NewBadRequest(gin.H{"Message": "Empty Body? May be!"})
		c.JSON(restResponse.Status, restResponse)
		return
	}

	if isValid != nil {
		log.Errorln(isValid)
		restResponse := response.NewBadRequest(gin.H{"Message": "Invalid Body"})
		c.JSON(restResponse.Status, restResponse)
		return
	}

	log.Infoln("Producing logs on destination", producerEntity.DestinationIP, "over port", producerEntity.DestinationPort)
	go producer_service.Produce(producerEntity, constant.ExecutionModeProduce)
	c.JSON(http.StatusAccepted, gin.H{"Message": "Execution started"})
}

func ProduceTest(c *gin.Context) {
	var producerEntity producer_model.ProducerEntity

	isValid := c.ShouldBindJSON(&producerEntity)

	if isValid == io.EOF {
		log.Errorln(isValid)
		restResponse := response.NewBadRequest(gin.H{"Message": "Empty Body? May be!"})
		c.JSON(restResponse.Status, restResponse)
		return
	}

	if isValid != nil {
		log.Errorln(isValid)
		response := response.NewBadRequest(gin.H{"Message": "Invalid Body"})
		c.JSON(response.Status, response)
		return
	}

	response := producer_service.Produce(producerEntity, constant.ExecutionModeTest)
	c.JSON(response.Status, response)
}
