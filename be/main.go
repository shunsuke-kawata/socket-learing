package main

import (
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

		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	// //デバイスを受け取る
	// router.PORT("/", func(c *gin.Context) {

	// 	c.JSON(200, gin.H{
	// 		"message": "success",
	// 	})
	// })

	//デバイス名に関してはタスクを追加するときについでに一緒に追加する

	//タスク一覧を受け取る
	router.GET("task",func(c *gin.Context)){

	}

	//タスクを追加する
	router.POST("task",func(c *gin.Context)){

	}
	
	//タスクの内容を更新する
	router.PUT("task/:id",func(c *gin.Context)){

	}

	//タスクを削除する
	router.DELETE("task/:id",func(c *gin.Context)){

	}

	return router
}

func main() {
	router := CreateRouter()
	router.Run(":" + os.Getenv("GO_PORT"))
}
