package networkUtils

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/yjagdale/siem-data-producer/Formatter"
	"github.com/yjagdale/siem-data-producer/utils/error_response"
	"net"
	"strings"
	"time"
)

func GetConnection(destinationServer string, protocol string) (net.Conn, error) {
	var conn net.Conn
	var err error
	switch strings.ToUpper(protocol) {
	case "TCP":
		log.Infoln("Building tcp connection")
		conn, err = net.DialTimeout("tcp", destinationServer, 40*time.Second)
		break
	case "UDP":
		log.Infoln("Building UDP connection")
		conn, err = net.Dial("udp", destinationServer)
		break
	}
	if err != nil {
		log.Errorln("could not connect to server: ", err)
		return nil, err
	}
	return conn, nil
}

func ProduceLogs(connection net.Conn, logs []string) *error_response.RestErr {
	success := 0
	failed := 0
	for _, logLine := range logs {
		logLine = Formatter.FormatLog(logLine)
		err := pushLog(connection, logLine)
		if err != nil {
			failed++
		} else {
			success++
		}
	}
	return error_response.NewPartialProcessError(success, failed)
}

func pushLog(connection net.Conn, logLine string) error {
	noOfBytes, err := fmt.Fprintln(connection, logLine)
	if err != nil {
		return err
	}

	log.Debugln("Published ", noOfBytes)
	return nil
}