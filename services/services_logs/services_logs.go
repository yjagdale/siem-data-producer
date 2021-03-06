package services_logs

import (
	"github.com/yjagdale/siem-data-producer/models/model_logs"
	"github.com/yjagdale/siem-data-producer/utils/utils_http/response"
)

func GetAvailableLogFiles(fileName string) (*[]model_logs.Logs, *response.RestResponse) {
	var logs model_logs.Logs

	if fileName != "" {
		logs.DeviceType = fileName
		return logs.Get()
	} else {
		return logs.GetAll()
	}

}
