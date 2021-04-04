package logs_model

import "gorm.io/gorm"

type Logs struct {
	gorm.Model
	DeviceType    string          `json:"device_type"`
	AvailableLogs []AvailableLogs `json:"available_logs"`
}

type AvailableLogs struct {
	gorm.Model
	LogsID   uint   `json:"log_id"`
	Location string `json:"location"`
}
