package model_logs

type ProducerEntity struct {
	Logs            []string `json:"logs" binding:"required"`
	DestinationIP   string   `json:"destination_ip" binding:"required"`
	DestinationPort int      `json:"destination_port" binding:"required"`
	Protocol        string   `json:"protocol" binding:"required"`
}

func (e ProducerEntity) ProduceSingle() {

}
