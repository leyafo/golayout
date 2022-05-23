package helper

import (
	"encoding/json"
	"fmt"
	"golayout/pkg/logger"
	"net/http"
)

type ErrorType uint

const (
	OK           ErrorType = 0
	UnknownError ErrorType = 1
	HTTPError    ErrorType = 2
)

var errMapping = map[ErrorType]string{
	OK:           "",
	UnknownError: "未知错误",
	HTTPError:    "HTTP错误",
}

type CommonResponse struct {
	Code         ErrorType   `json:"code,omitempty"`
	Message      string      `json:"message"`
	ExtraMessage string      `json:"extra_message,omitempty"`
	Data         interface{} `json:"data"`
}

func writeResponse(w http.ResponseWriter, cr *CommonResponse, statusCode int) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(cr); err != nil {
		logger.Error(err)
	}
}

func ResponseWithUserMsg(w http.ResponseWriter, httpStatusCode int,
	data interface{}, msgCode ErrorType, message string) {
	localMsg := errMapping[msgCode]
	cr := CommonResponse{
		Code:         msgCode,
		Message:      localMsg,
		ExtraMessage: message,
		Data:         data,
	}
	writeResponse(w, &cr, httpStatusCode)
}

func ResponseWithError(w http.ResponseWriter, httpStatusCode int, msgCode ErrorType, err error) {
	msg, ok := errMapping[msgCode]
	if !ok {
		msg = errMapping[1]
	}
	msg = fmt.Sprintf("%s: %s", msg, err.Error())

	cr := CommonResponse{
		Code:    msgCode,
		Message: msg,
	}
	writeResponse(w, &cr, httpStatusCode)
}

func ResponseFailureMsg(w http.ResponseWriter, httpStatusCode int, msg string) {
	cr := CommonResponse{
		Code:    HTTPError,
		Message: msg,
	}
	writeResponse(w, &cr, httpStatusCode)
}
func ResponseFailureError(w http.ResponseWriter, httpStatusCode int, err error) {
	cr := CommonResponse{
		Code:    HTTPError,
		Message: err.Error(),
	}
	writeResponse(w, &cr, httpStatusCode)
}

func Response(
	w http.ResponseWriter, httpStatusCode int,
	data interface{}, MsgCode ErrorType) {
	if httpStatusCode == http.StatusOK {
		MsgCode = 0
	}

	msg, ok := errMapping[MsgCode]
	if !ok {
		msg = errMapping[UnknownError]
	}
	cr := CommonResponse{
		Code:    MsgCode,
		Message: msg,
		Data:    data,
	}
	writeResponse(w, &cr, httpStatusCode)
}

func ResponseSuccessData(w http.ResponseWriter, data interface{}) {
	cr := CommonResponse{
		Code:    OK,
		Message: "",
		Data:    data,
	}
	writeResponse(w, &cr, http.StatusOK)
}
