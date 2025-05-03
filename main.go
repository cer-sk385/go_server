// Package main は、Ginフレームワークを使用したシンプルなRESTful APIサーバーを提供します。
// 基本的なGET/POSTエンドポイントを実装し、JSONレスポンスを返します。
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response はAPIレスポンスの基本構造を定義します。
type Response struct {
	Message string `json:"message"`
}

// main はサーバーを初期化し、ルーティングを設定して起動します。
func main() {
	// Ginルーターの初期化
	r := gin.Default()

	// ルーティングの設定
	r.GET("/", HandleHome)
	r.GET("/hello", HandleHello)
	r.GET("/goodbye", HandleGoodbye)

	// POSTの例
	r.POST("/post", HandlePost)

	// サーバー起動
	r.Run(":8080")
}

// HandleHome はルートパスへのGETリクエストを処理します。
func HandleHome(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Message: "Hello, World!",
	})
}

// HandleHello は/helloエンドポイントへのGETリクエストを処理します。
func HandleHello(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Message: "こんにちは！",
	})
}

// HandleGoodbye は/goodbyeエンドポイントへのGETリクエストを処理します。
func HandleGoodbye(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Message: "さようなら！",
	})
}

// HandlePost は/postエンドポイントへのPOSTリクエストを処理します。
// JSONボディからメッセージを受け取り、そのメッセージを返します。
func HandlePost(c *gin.Context) {
	var json struct {
		Message string `json:"message"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Message: "Invalid JSON format",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"received": json.Message,
	})
}
