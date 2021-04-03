package error_response

import (
	"net/http"
	"strconv"
)

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewBadRequest(message string) *RestErr {
	return &RestErr{Message: message, Status: http.StatusBadRequest}
}

func NewPartialProcessError(successCount int, failCount int) *RestErr {
	return &RestErr{Message: "success:" + strconv.Itoa(successCount) + " failed: " + strconv.Itoa(failCount), Status: http.StatusPartialContent}
}

func NewNotFound(message string) *RestErr {
	return &RestErr{Message: message, Status: http.StatusNotFound}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{Message: message, Status: http.StatusInternalServerError}
}

func NewNotImplementedError(message string) *RestErr {
	return &RestErr{Message: message, Status: http.StatusNotImplemented}
}

func NewEntityAlreadyExists(message string) *RestErr {
	return &RestErr{Message: message, Status: http.StatusAlreadyReported}
}

func NewAuthError(message string) *RestErr {
	return &RestErr{Message: message, Status: http.StatusUnauthorized}

}

func NewCollectorTypeError(message string) *RestErr {
	return &RestErr{Message: message, Status: http.StatusBadRequest}

}
