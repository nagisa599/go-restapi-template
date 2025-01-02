package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
	openapi "github.com/nasunagisa/restapi/app/gen"
	"github.com/nasunagisa/restapi/app/internal/usecase"
)

type IUserHandler interface {
	// (GET /users)
	GetUserList(ctx echo.Context, params openapi.GetUserListParams) error
	// (GET /users/{userId})
	GetUser(ctx echo.Context, userId int64) error
}

type userHandler struct {
	uu usecase.IUserUsecase
}


func NewUserHandler(uu usecase.IUserUsecase) IUserHandler {
	return &userHandler{uu}
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
	user, err := uh.uu.GetUser(userId)
	if err != nil {
		return ctx.JSON(500, err)
	}
	userRes := &openapi.User{
		Name: user.Name,
	}

	fmt.Println(userRes)
	return ctx.JSON(200, userRes)
}