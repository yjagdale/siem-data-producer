package file_producer_service

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/yjagdale/siem-data-producer/models/file_producer_model"
	"github.com/yjagdale/siem-data-producer/utils/networkUtils"
	"github.com/yjagdale/siem-data-producer/utils/response"
	"net"
	"os"
	"strconv"
)

func PublishFile(publisher file_producer_model.FileProducer) *response.RestErr {
	stats, err := os.Stat(publisher.Path)
	if err != nil {
		return &response.RestErr{Status: 400, Message: gin.H{"Message": "File does not exists"}}
	}

	destinationServer := publisher.DestinationIP + ":" + strconv.Itoa(publisher.DestinationPort)
	connection, err := networkUtils.GetConnection(destinationServer, publisher.Protocol)
	if err != nil {
		return response.NewBadRequest(gin.H{"error": err.Error()})
	}
	defer func(connection net.Conn) {
		err := connection.Close()
		if err != nil {
			log.Infoln("Error while closing file")
		}
	}(connection)

	log.Infoln("File existing and processing file. Records available in file are", stats.Size())
	publisher.ReadAndPublish(connection)
	return nil
}
