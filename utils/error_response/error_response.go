package error_response

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewBadRequest(message string) *RestErr {
	return &RestErr{Message: message, Status: http.StatusBadRequest}
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
