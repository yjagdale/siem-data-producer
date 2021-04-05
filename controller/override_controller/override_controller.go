package override_controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/yjagdale/siem-data-producer/models/override_model"
	"github.com/yjagdale/siem-data-producer/services/override_service"
	"github.com/yjagdale/siem-data-producer/utils/response"
	"io"
)

func AddOverride(c *gin.Context) {

	var overrideEntity []override_model.Override

	isValid := c.ShouldBindJSON(&overrideEntity)

	if isValid == io.EOF {
		log.Errorln(isValid)
		response := response.NewBadRequest(gin.H{"Message": "Empty Body? May be!"})
		c.JSON(response.Status, response)
		return
	}

	if isValid != nil {
		log.Errorln(isValid)
		response := response.NewBadRequest(gin.H{"Message": "Invalid Body"})
		c.JSON(response.Status, response)
		return
	}

	log.Infoln("Adding Override values")

	repo := override_service.StoreOverrideValues(overrideEntity)
	c.JSON(repo.Status, repo.Message)

}
