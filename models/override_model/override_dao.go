package override_model

type Override struct {
	Key    string   `json:"key" binding:"required"`
	Values []string `json:"values" binding:"required"`
}
