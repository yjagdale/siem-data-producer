package files_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yjagdale/siem-data-producer/services/file_upload_service"
)

func Upload(c *gin.Context) {
	file_upload_service.UploadFile(c)
}
