package appcode

import "fmt"

type AppError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func New(code int, msg string) *AppError {
	return &AppError{code, msg}
}

func (a AppError) Error() string {
	return fmt.Sprintf("code %d msg:%s", a.Code, a.Msg)
}
