package producer_model

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/yjagdale/siem-data-producer/Formatter"
	"github.com/yjagdale/siem-data-producer/utils/networkUtils"
	"github.com/yjagdale/siem-data-producer/utils/response"
	"net/http"
	"strconv"
)

type ProducerEntity struct {
	Logs            []string `json:"logs" binding:"required"`
	DestinationIP   string   `json:"destination_ip" binding:"required"`
	DestinationPort int      `json:"destination_port" binding:"required"`
	Protocol        string   `json:"protocol" binding:"required"`
	Iterations      int      `json:"iterations" binding:"required"`
}

func (entity ProducerEntity) Produce() *response.RestErr {
	restResponse := response.RestErr{}
	destinationServer := entity.DestinationIP + ":" + strconv.Itoa(entity.DestinationPort)
	connection, err := networkUtils.GetConnection(destinationServer, entity.Protocol)
	if err != nil {
		return response.NewBadRequest(gin.H{"error": err.Error()})
	}
	defer connection.Close()
	var execution map[string]gin.H
	execution = make(map[string]gin.H)
	for i := 0; i < entity.Iterations; i++ {
		runStatus := networkUtils.ProduceLogs(i, connection, entity.Logs)
		execution["iteration_"+strconv.Itoa(i)] = runStatus
	}
	restResponse.Message = gin.H{"Execution Status": execution}
	return &restResponse
}

func (entity ProducerEntity) ProduceTest() *response.RestErr {
	outputLogs := response.RestErr{}
	outputLogs.Status = http.StatusOK
	log.Infoln("Debugging logs")
	var formattedLogs []string
	for _, logLine := range entity.Logs {
		formattedLogs = append(formattedLogs, Formatter.FormatLog(logLine))
	}
	outputLogs.Message = gin.H{"Formatted Logs": formattedLogs}
	return &outputLogs
}
