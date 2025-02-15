package domain

import (
	"fmt"

	"go.uber.org/zap"
)


type MyError struct {
  Code       int
  Msg        string
  StackTrace string
}

// Error error interfaceを実装
func (me *MyError) Error() string {
  return fmt.Sprintf("my error: code[%d], message[%s]", me.Code, me.Msg)
}

// New コンストラクタ
func New(code int, msg string) *MyError {
  stack := zap.Stack("").String
  return &MyError{
    Code:       code,
    Msg:        msg,
    StackTrace: stack,
  }
}
func NewInternalServerError() *MyError {
	return New(500, "Internal Server Error")
}

func NewForbiddenError() *MyError {
	return New(403, "Forbidden")
}