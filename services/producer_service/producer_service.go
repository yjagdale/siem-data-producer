package producer_service

import (
	"github.com/yjagdale/siem-data-producer/models/producer_model"
	"github.com/yjagdale/siem-data-producer/utils/error_response"
)

func Produce(producerEntity producer_model.ProducerEntity) []*error_response.RestErr {

	if len(producerEntity.Logs) <= 0 {
		var errorResp []*error_response.RestErr
		errorResp = append(errorResp, error_response.NewBadRequest("Log Lines Cant be 0"))
		return errorResp
	}
	return producerEntity.Produce()
}
