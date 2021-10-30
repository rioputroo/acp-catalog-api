package response

import "net/http"

type SuccessResponseCode string

//List of success response status
const (
	Success SuccessResponseCode = "success"
)

//SuccessResponse default payload response
type SuccessResponse struct {
	Code    SuccessResponseCode `json:"code"`
	Message string              `json:"message"`
	Data    interface{}         `json:"data"`
}

//NewSuccessResponse create new success payload
func NewSuccessResponse(data interface{}, msg string) (int, SuccessResponse) {
	return http.StatusOK, SuccessResponse{
		Success,
		msg,
		data,
	}
}

//NewSuccessResponse create new success payload
func NewSuccessResponseWithoutData(msg string) (int, SuccessResponse) {
	return http.StatusOK, SuccessResponse{
		Success,
		msg,
		map[string]interface{}{},
	}
}

//NewSuccessResponse create new success payload
func NewSuccessResponseNoContent(msg string) (int, SuccessResponse) {
	return http.StatusNoContent, SuccessResponse{
		Success,
		msg,
		map[string]interface{}{},
	}
}
