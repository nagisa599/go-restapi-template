package main

import (
	"fmt"
	"net/http"
)

func main() {
    // ハンドラ関数の定義
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, world!")
    })

    // サーバーをポート8080で起動
    fmt.Println("Server is running on http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server: ", err)
    }
}
