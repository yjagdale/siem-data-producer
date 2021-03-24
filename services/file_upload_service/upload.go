package file_upload_service

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/yjagdale/siem-data-producer/config/constant"
)

type FileService interface {
	UploadFile(c *gin.Context)
}

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
	url := c.Request.Header.Get("origin")
	fmt.Println(url)
	err = os.MkdirAll(constant.FileOutputBasePath+deviceType+deviceVendor, 0755)
	if err != nil {
		log.Errorln("Error while creating dir", err.Error())
	}
	if err := c.SaveUploadedFile(file, constant.FileOutputBasePath+deviceType+deviceVendor+filename); err != nil {
		c.Redirect(302, url+"/ui/?uploadStatus="+fmt.Sprintf("Upload failed due to %s ", err))
		return
	}
	c.Redirect(302, url+"/ui/?uploadStatus="+fmt.Sprintf("File %s uploaded successfully", file.Filename))
}
