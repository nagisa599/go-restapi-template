package main

import (
	router "github.com/nasunagisa/restapi/app/infrastructure"
	"github.com/nasunagisa/restapi/app/internal/handler"
)

func main() {
    // ハンドラ関数の定義
    userHandler := handler.NewUserHandler()
    todoHandler := handler.NewTodoHandler()

    e := router.NesRouter(userHandler, todoHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
