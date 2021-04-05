package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewOk(message gin.H) *RestErr {
	return &RestErr{Message: message, Status: http.StatusOK}
}
