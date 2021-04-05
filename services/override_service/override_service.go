package override_service

import (
	"github.com/gin-gonic/gin"
	"github.com/yjagdale/siem-data-producer/models/override_model"
	"github.com/yjagdale/siem-data-producer/utils/response"
	"strings"
)

func StoreOverrideValues(overrideEntities []override_model.Override) *response.RestErr {
	responseEntity := response.RestErr{}
	responseMessage := make(map[string]string)
	for _, overrideEntity := range overrideEntities {
		if len(overrideEntity.Values) == 0 {
			responseMessage[overrideEntity.Key] = "Please provide values"
			continue
		}

		if strings.TrimSpace(overrideEntity.Key) == "" {
			responseMessage[overrideEntity.Key] = "Key should be non empty"
			continue
		}
		responseMessage[overrideEntity.Key] = overrideEntity.Save()
	}
	responseEntity.Message = gin.H{"Store Response": responseMessage}
	return &responseEntity

}
