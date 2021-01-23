package file_upload_service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/yjagdale/siem-data-producer/config/constant"
	"net/http"
	"os"
	"path/filepath"
)

func UploadFile(c *gin.Context) {
	deviceType := c.PostForm("deviceType") + "/"
	if deviceType == "/" {
		c.String(http.StatusBadRequest, "Device type is mandatory")
		return
	}
	deviceVendor := c.PostForm("deviceVendor") + "/"

	if deviceVendor == "/" {
		c.String(http.StatusBadRequest, "Device vendor is mandatory")
		return
	}
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	filename := filepath.Base(file.Filename)
	fileExtension := filepath.Ext(filename)
	if fileExtension != ".log" {
		c.String(http.StatusBadRequest, fmt.Sprintf("%s is not supported at the moment. Supported formats .log", fileExtension))
		return
	}

	err = os.MkdirAll("/tmp/storage/logs/"+deviceType+deviceVendor, 0755)
	if err != nil {
		log.Errorln("Error while creating dir", err.Error())
	}
	if err := c.SaveUploadedFile(file, constant.FileOutputBasePath+deviceType+deviceVendor+filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully", file.Filename))
}
