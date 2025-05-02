package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// ルートパスへのハンドラを設定
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "こんにちは！")
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "さようなら！")
	})

	// ポート番号を指定してサーバーを起動
	fmt.Println("サーバーを起動します: http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
