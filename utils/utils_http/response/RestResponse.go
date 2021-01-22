package response

import "net/http"

type RestResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func DBConnectError(message string) *RestResponse {
	var restResponse RestResponse

	restResponse.Code = http.StatusInternalServerError
	restResponse.Message = message

	return &restResponse
}

func NoRecordsFound() *RestResponse {
	var restResponse RestResponse

	restResponse.Code = http.StatusNotFound
	restResponse.Message = "No records found"

	return &restResponse
}
