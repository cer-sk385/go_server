package main

import (
	"fmt"
	"io/ioutil"
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

	// GETリクエスト
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "GETメソッドのみ許可されています", http.StatusMethodNotAllowed)
			return
		}
		fmt.Fprintf(w, "GETリクエストを受け取りました！")
	})

	// POSTリクエスト
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "POSTメソッドのみ許可されています", http.StatusMethodNotAllowed)
			return
		}
		// POSTデータを読み取る
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "リクエストボディの読み取りに失敗しました", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "POSTリクエストを受け取りました！データ: %s", string(body))
	})

	// ポート番号を指定してサーバーを起動
	fmt.Println("サーバーを起動します: http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
