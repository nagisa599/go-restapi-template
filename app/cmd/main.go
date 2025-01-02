package main

import (
	router "github.com/nasunagisa/restapi/app/infrastructure"
	"github.com/nasunagisa/restapi/app/internal/handler"
	"github.com/nasunagisa/restapi/app/internal/usecase"
)

func main() {
    // ハンドラ関数の定義
    userUsecase := usecase.NewUserUsecase()
    userHandler := handler.NewUserHandler(userUsecase)
    todoHandler := handler.NewTodoHandler()

    e := router.NesRouter(userHandler, todoHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
