package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Gin のデフォルトのルーターを作成
	r := gin.Default()

	// ルートパス（"/"）にアクセスした時に index.html を表示
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "簡易Twitter",
		})
	})

	// HTML テンプレートをロード
	r.LoadHTMLGlob("templates/*")

	// サーバーを起動
	r.Run(":8080") // ポート8080で起動
}
