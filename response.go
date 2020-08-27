package fin

const (
	StatusFail    string = "fail"
	StatusSuccess string = "success"
)

const (
	MessageOk string = "OK"
)

type BaseSuccessResponse struct {
	Status  string      `json:"status" example:"success"`
	Message string      `json:"message" example:"OK"`
	Data    interface{} `json:"data"`
}

func NewSuccessResponse(data interface{}) BaseSuccessResponse {
	r := BaseSuccessResponse{
		Status:  StatusSuccess,
		Message: MessageOk,
		Data:    data,
	}
	return r
}

type ErrorResponse struct {
	Status  string `json:"status" example:"fail"`
	Message string `json:"message" example:"Error message will be show here"`
}

func NewErrorResponse(message string) ErrorResponse {
	errorResponse := ErrorResponse{}
	errorResponse.Status = StatusFail
	errorResponse.Message = message
	return errorResponse
}
