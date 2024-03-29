package file_producer_model

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/yjagdale/siem-data-producer/utils/files"
	"github.com/yjagdale/siem-data-producer/utils/networkUtils"
	"github.com/yjagdale/siem-data-producer/utils/response"
	"net"
	"os"
	"strconv"
)

type FileProducer struct {
	Path            string `json:"path"`
	DestinationIP   string `json:"destination_ip"`
	DestinationPort int    `json:"destination_port"`
	Protocol        string `json:"protocol"`
	Eps             []int  `json:"EPS" binding:"required"`
}

func (publisher FileProducer) ReadAndPublish(connection net.Conn) *response.RestErr {
	restResponse := response.RestErr{}

	fi, err := os.Stat(publisher.Path)
	if err != nil {
		return &response.RestErr{Status: 400, Message: gin.H{"Message": "File is empty"}}
	}
	// get the size
	size := fi.Size()

	if size > 1594682 {
		log.Infoln("Large file, Producer will run in background")
		files.ReadAndPublishInChunk(publisher.Path, connection)
		restResponse.Message = gin.H{"Message": "Large file, Execution started"}
		restResponse.Status = 201
		return &restResponse
	} else {
		logLines := files.ReadFileLineByLine(publisher.Path)
		log.Infoln("File has ", len(logLines), "Records")
		if len(logLines) <= 0 {
			return &response.RestErr{Status: 400, Message: gin.H{"Message": "File is empty"}}
		}
		var execution map[string]gin.H
		execution = make(map[string]gin.H)
		for i := 0; i < 2; i++ {
			runStatus := networkUtils.ProduceLogs(i, connection, logLines)
			execution["iteration_"+strconv.Itoa(i)] = runStatus
		}
		restResponse.Message = gin.H{"Execution Status": execution}
		return &restResponse
	}
}
