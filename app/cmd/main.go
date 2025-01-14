package main

import (
	"github.com/nasunagisa/restapi/app/infrastructure"
	"github.com/nasunagisa/restapi/app/internal/domain/repository"
	"github.com/nasunagisa/restapi/app/internal/handler"
	"github.com/nasunagisa/restapi/app/internal/usecase"
)

func main() {
    // ハンドラ関数の定義
    db := infrastructure.NewDatabase()
    userRepository := repository.NewUserRepository(db)
    userUsecase := usecase.NewUserUsecase(userRepository)
    
    userHandler := handler.NewUserHandler(userUsecase)
    todoHandler := handler.NewTodoHandler()
    errorHandler := handler.NewErrorHandler()

    e := infrastructure.NesRouter(userHandler, todoHandler, errorHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
