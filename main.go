package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Ginルーターの初期化
	r := gin.Default()

	// ルーティングの設定
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "こんにちは！",
		})
	})

	r.GET("/goodbye", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "さようなら！",
		})
	})

	// POSTの例
	r.POST("/post", func(c *gin.Context) {
		var json struct {
			Message string `json:"message"`
		}

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"received": json.Message,
		})
	})

	// サーバー起動
	r.Run(":8080")
}
