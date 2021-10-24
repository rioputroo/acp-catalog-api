package response

import "net/http"

type errorControllerResponseCode string

const (
	ErrBadRequest errorControllerResponseCode = "Bad request"
	errForbidden  errorControllerResponseCode = "Forbidden"
)

type ControllerResponse struct {
	Code    errorControllerResponseCode `json:"code"`
	Message string                      `json:"message"`
	Data    interface{}                 `json:"data"`
}

//NewBadRequestResponse bad request format response
func NewBadRequestResponse() (int, ControllerResponse) {
	return http.StatusBadRequest, ControllerResponse{
		ErrBadRequest,
		"Bad request",
		map[string]interface{}{},
	}
}

//NewForbiddenResponse default for Forbidden error response
func NewForbiddenResponse() (int, ControllerResponse) {
	return http.StatusForbidden, ControllerResponse{
		errForbidden,
		"Forbidden",
		map[string]interface{}{},
	}
}
