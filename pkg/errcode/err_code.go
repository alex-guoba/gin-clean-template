package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	// 错误码
	Code int `json:"code"`
	// 错误消息
	Msg string `json:"msg"`
	// 详细信息
	Details []string `json:"details"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("error code %d already exist", code))
	}
	codes[code] = msg
	return &Error{Code: code, Msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("error code %d, detail: %s", e.Code, e.Msg)
}

// func (e *Error) Code() int {
// 	return e.code
// }

// func (e *Error) Msg() string {
// 	return e.msg
// }

func (e *Error) Msgf(args []any) string {
	return fmt.Sprintf(e.Msg, args...)
}

// func (e *Error) Details() []string {
// 	return e.details
// }

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.Details = []string{}
	newError.Details = append(newError.Details, details...)

	return &newError
}

func (e *Error) StatusCode() int {
	switch e.Code {
	case Success.Code:
		return http.StatusOK
	case ServerError.Code:
		return http.StatusInternalServerError
	case InvalidParams.Code:
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code:
		fallthrough
	case UnauthorizedTokenError.Code:
		fallthrough
	case UnauthorizedTokenGenerate.Code:
		fallthrough
	case UnauthorizedTokenTimeout.Code:
		return http.StatusUnauthorized
	case TooManyRequests.Code:
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}
