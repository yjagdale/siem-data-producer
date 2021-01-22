package model_logs

import (
	"errors"
	"github.com/yjagdale/siem-data-producer/database"
	"github.com/yjagdale/siem-data-producer/utils/utils_http/response"
	"gorm.io/gorm"
)

func (log *Logs) GetAll() (*[]Logs, *response.RestResponse) {
	var logs []Logs
	if database.ValidateDBConnection() {
		err := database.Connection.Debug().Find(&logs).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NoRecordsFound()
		} else {
			return &logs, nil
		}
	} else {
		return nil, response.DBConnectError("connection to database failed")
	}

}
