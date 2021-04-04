package producer_model

import (
	"github.com/yjagdale/siem-data-producer/utils/error_response"
	"github.com/yjagdale/siem-data-producer/utils/networkUtils"
	"strconv"
)

type ProducerEntity struct {
	Logs            []string `json:"logs" binding:"required"`
	DestinationIP   string   `json:"destination_ip" binding:"required"`
	DestinationPort int      `json:"destination_port" binding:"required"`
	Protocol        string   `json:"protocol" binding:"required"`
	Iterations      int      `json:"iterations" binding:"required"`
}

func (entity ProducerEntity) Produce() []*error_response.RestErr {
	var response []*error_response.RestErr
	destinationServer := entity.DestinationIP + ":" + strconv.Itoa(entity.DestinationPort)
	connection, err := networkUtils.GetConnection(destinationServer, entity.Protocol)
	if err != nil {
		response := append(response, error_response.NewBadRequest(err.Error()))
		return response
	}
	defer connection.Close()
	for i := 0; i < entity.Iterations; i++ {
		response = append(response, networkUtils.ProduceLogs(connection, entity.Logs))
	}
	return response
}
