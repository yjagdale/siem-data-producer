package producer_service

import (
	log "github.com/sirupsen/logrus"
	"github.com/yjagdale/siem-data-producer/models/model_logs"
	"github.com/yjagdale/siem-data-producer/utils/error_response"
	"github.com/yjagdale/siem-data-producer/utils/tcpUtils"
)

func ProduceSingle(producerEntity model_logs.ProducerEntity) *error_response.RestErr {

	if len(producerEntity.Logs) <= 0 {
		return error_response.NewBadRequest("Log Lines Cant be 0")
	}

	if producerEntity.Protocol == "tcp" {
		log.Infoln("Destination is TCP, Validating connection")
		err := tcpUtils.ValidateConnection(producerEntity.DestinationIP, producerEntity.DestinationPort)
		if err != nil {
			return err
		}
	} else {
		log.Infoln("Pushing logs on udp. No validation needed.")
		return nil
	}
	return nil
}
