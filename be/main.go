package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// ルーターインスタンスを作成
func CreateRouter() *gin.Engine {
	// routerのインスタンスを作成
	router := gin.Default()

	// corsの設定
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
	}))

	router.GET("/", func(c *gin.Context) {
		// バックエンドログを標準出力
		fmt.Println("Request received at /")

		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	return router
}

func main() {
	router := CreateRouter()
	router.Run(":" + os.Getenv("GO_PORT"))
}
