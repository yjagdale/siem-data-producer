package tcpUtils

import (
	log "github.com/sirupsen/logrus"
	"github.com/yjagdale/siem-data-producer/utils/error_response"
	"net"
	"strconv"
)

func ValidateConnection(destinationIP string, destinationPort int) *error_response.RestErr {
	destinationServer := destinationIP + ":" + strconv.Itoa(destinationPort)
	err := validateTCP(destinationServer)

	if err != nil {
		return error_response.NewBadRequest("Destination IP on Port is not listening. TCP ack failed")
	}
	return nil
}

func validateTCP(destinationIp string) error {
	conn, err := net.Dial("tcp", destinationIp)
	if err != nil {
		log.Errorln("could not connect to TCP server: ", err)
		return err
	}
	defer conn.Close()
	return nil
}
