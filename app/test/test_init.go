package test

import (
	"github.com/nasunagisa/restapi/app/infrastructure"
	"github.com/nasunagisa/restapi/app/internal/domain/repository"
	"github.com/nasunagisa/restapi/app/internal/handler"
	"github.com/nasunagisa/restapi/app/internal/usecase"
)

// initTest initializes the test environment and returns an IUserHandler.
func InitTest() handler.IUserHandler {
	db := infrastructure.NewTestDatabase()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)
	return userHandler
}