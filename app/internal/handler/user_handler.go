package handler

import (
	"github.com/labstack/echo/v4"
	openapi "github.com/nasunagisa/restapi/app/gen"
)

type IUserHandler interface {
	// (GET /users)
	GetUserList(ctx echo.Context, params openapi.GetUserListParams) error
		// ユーザー情報取得
	// (GET /users/{userId})
	GetUser(ctx echo.Context, userId int64) error
}

type userHandler struct {}


func NewUserHandler() IUserHandler {
	return &userHandler{}
}


func (uh *userHandler) GetUserList(ctx echo.Context, params openapi.GetUserListParams ) error {
	userRes :=
		 []*openapi.User{
			{
				UserId: 1,
				Name: "user1",
			},
			{
				UserId: 2,
				Name: "user2",
			},
		}

	return ctx.JSON(200, userRes)
}

func (uh *userHandler) GetUser(ctx echo.Context, userId int64) error {
	userRes := &openapi.User{
		UserId: userId,
		Name: "user1",
	}

	return ctx.JSON(200, userRes)
}