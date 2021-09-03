package producer_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/yjagdale/siem-data-producer/config/constant"
	"github.com/yjagdale/siem-data-producer/models/logs_model"
	"github.com/yjagdale/siem-data-producer/models/producer_model"
	"github.com/yjagdale/siem-data-producer/services/producer_service"
	"github.com/yjagdale/siem-data-producer/utils/response"
	"io"
	"net/http"
	"time"
)

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
	if logs_model.ContinuesExecution == nil {
		logs_model.ContinuesExecution = make(map[string]bool)
		logs_model.ContinuesExecution["status"] = true
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
	logs_model.ContinuesExecution[executionId] = true
	go func() {
		for logs_model.ContinuesExecution[executionId] != false {
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

	if logs_model.ContinuesExecution == nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "No Execution going on"})
		return
	}

	if logs_model.ContinuesExecution[executionID] != true {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Execution already stopped"})
		return
	}

	logs_model.ContinuesExecution[executionID] = false
	c.JSON(http.StatusOK, gin.H{"Message": "Execution stopped. It will take ~5  mins to kill"})
	return
}

func GetExecutions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Executions": logs_model.ContinuesExecution})
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
