package override_model

import (
	"gorm.io/gorm"
)

var OverrideValues map[string][]string

type OverrideConfig struct {
	gorm.Model
	Key    string   `json:"key" binding:"required" gorm:"primaryKey"`
	Values []string `json:"values" binding:"required" gorm:"type:text[]"`
}

func connectDB() {
	if OverrideValues == nil {
		OverrideValues = make(map[string][]string)
	}
	return
}

func (overrides Override) Save() string {
	connectDB()
	if _, ok := OverrideValues[overrides.Key]; ok {
		OverrideValues[overrides.Key] = append(OverrideValues[overrides.Key], overrides.Values...)
	} else {
		OverrideValues[overrides.Key] = overrides.Values
	}

	return "Override Added Successfully"
}
