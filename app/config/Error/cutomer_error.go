package Error

import (
	"fmt"

	"go.uber.org/zap"
)


type CustomerError struct {
  Code       int
  Msg        string
  StackTrace string
}

// Error error interfaceを実装
func (me *CustomerError) Error() string {
  return fmt.Sprintf("my error: code[%d], message[%s]", me.Code, me.Msg)
}

// New コンストラクタ
func New(code int, msg string) *CustomerError {
  stack := zap.Stack("").String
  return &CustomerError{
    Code:       code,
    Msg:        msg,
    StackTrace: stack,
  }
}
func NewInternalServerError() *CustomerError {
	return New(500, "Internal Server Error")
}

func NewForbiddenError() *CustomerError {
	return New(403, "Forbidden")
}

func NewNotFoundError() *CustomerError {
  return New(404, "Not Found")
}