package producer_service

import (
	"github.com/gin-gonic/gin"
	"github.com/yjagdale/siem-data-producer/config/constant"
	"github.com/yjagdale/siem-data-producer/models/producer_model"
	"github.com/yjagdale/siem-data-producer/utils/response"
)

func Produce(producerEntity producer_model.ProducerEntity, executionMode string) *response.RestErr {

	if len(producerEntity.Logs) <= 0 {
		return response.NewBadRequest(gin.H{"Message": "Log Lines Cant be 0"})
	}
	if executionMode == constant.ExecutionModeProduce {
		return producerEntity.Produce()
	} else if executionMode == constant.ExecutionModeTest {
		return producerEntity.ProduceTest()
	}
	return nil
}
