package producer_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/yjagdale/siem-data-producer/config/constant"
	"github.com/yjagdale/siem-data-producer/models/producer_model"
	"github.com/yjagdale/siem-data-producer/services/producer_service"
	"github.com/yjagdale/siem-data-producer/utils/response"
	"io"
	"net/http"
	"time"
)

var continuesExecution map[string]bool

func Produce(c *gin.Context) {

	var producerEntity producer_model.ProducerEntity

	isValid := c.ShouldBindJSON(&producerEntity)

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

	log.Infoln("Producing logs on destination", producerEntity.DestinationIP, "over port", producerEntity.DestinationPort)
	resp := producer_service.Produce(producerEntity, constant.ExecutionModeProduce)
	c.JSON(resp.Status, resp)
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
func ProduceContinues(c *gin.Context) {
	if continuesExecution["status"] {
		continuesExecution = make(map[string]bool)
		continuesExecution["status"] = true
	}
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

	executionId := uuid.New().String()
	continuesExecution[executionId] = true
	go func() {
		for continuesExecution[executionId] != false {
			producer_service.Produce(producerEntity, constant.ExecutionModeProduce)
			time.Sleep(5 * time.Second)
		}
	}()
	c.JSON(http.StatusAccepted, gin.H{"Message": "Execution started id is: " + executionId})
}

func StopExecutor(c *gin.Context) {
	executionID := c.Param("executionId")

	if executionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Execution id needed"})
		return
	}

	if continuesExecution[executionID] != true {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Execution already stopped"})
		return
	}

	continuesExecution[executionID] = false
	c.JSON(http.StatusOK, gin.H{"Message": "Execution stopped. It will take ~5  mins to kill"})
	return
}

func GetExecutions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Executions": continuesExecution})
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
		resp := response.NewBadRequest(gin.H{"Message": "Invalid Body"})
		c.JSON(resp.Status, resp)
		return
	}

	resp := producer_service.Produce(producerEntity, constant.ExecutionModeTest)
	c.JSON(resp.Status, resp)
}
