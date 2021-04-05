package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RestErr struct {
	Message gin.H
	Status  int `json:"status"`
}

func NewBadRequest(message gin.H) *RestErr {
	return &RestErr{Message: message, Status: http.StatusBadRequest}
}

func NewPartialProcessError(message gin.H) *RestErr {
	return &RestErr{Message: message, Status: http.StatusPartialContent}
}

func NewNotFound(message gin.H) *RestErr {
	return &RestErr{Message: message, Status: http.StatusNotFound}
}

func NewInternalServerError(message gin.H) *RestErr {
	return &RestErr{Message: message, Status: http.StatusInternalServerError}
}

func NewNotImplementedError(message gin.H) *RestErr {
	return &RestErr{Message: message, Status: http.StatusNotImplemented}
}

func NewEntityAlreadyExists(message gin.H) *RestErr {
	return &RestErr{Message: message, Status: http.StatusAlreadyReported}
}

func NewAuthError(message gin.H) *RestErr {
	return &RestErr{Message: message, Status: http.StatusUnauthorized}

}

func NewCollectorTypeError(message gin.H) *RestErr {
	return &RestErr{Message: message, Status: http.StatusBadRequest}

}
