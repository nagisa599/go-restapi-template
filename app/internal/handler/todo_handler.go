package handler

import (
	"github.com/labstack/echo/v4"
	openapi "github.com/nasunagisa/restapi/app/gen"
)

type ITodoHandler interface {
	// todoリストの一覧を取得
	// (GET /todos/{userId})
	GetTodos(ctx echo.Context, userId int64) error
}

type todoHandler struct{}

func NewTodoHandler() ITodoHandler {
	return &todoHandler{}
}


func (th *todoHandler) GetTodos(ctx echo.Context, userId int64) error {
	todoRes :=
		[]*openapi.Todo{
			{
				Title:   "title1",
				Content: "content1",
			},
			{
				Title:   "title2",
				Content: "content2",
			},
		}

	return ctx.JSON(200, todoRes)
}
