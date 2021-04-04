package services_logs

import (
	"github.com/yjagdale/siem-data-producer/models/model_logs"
	"github.com/yjagdale/siem-data-producer/utils/utils_http/response"
)

func GetAvailableLogFiles(fileName string) (*[]logs_model.Logs, *response.RestResponse) {
	var logs logs_model.Logs

	if fileName != "" {
		logs.DeviceType = fileName
		return logs.Get()
	} else {
		return logs.GetAll()
	}

}
