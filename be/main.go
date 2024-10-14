package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shunsuke-kawata/socket-learning/model"
)

// データベースのスキーマ
type StatusParam struct {
	StatusName string
}
type AddressParam struct {
	IPAddress string
}

type TaskParam struct {
	Title       string
	Description string
	ColorCode   string
	IPAddress   string
}

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

	//GET ステータス一覧
	router.GET("/status", func(c *gin.Context) {
		statuses, err := model.ReadStatus()
		if err != nil {
			fmt.Println(err)
		} else {
			c.JSON(200, statuses)
			fmt.Printf("%T\n", statuses)
		}

	})
	// GET 接続記録のあるIPアドレス一覧
	router.GET("/address", func(c *gin.Context) {
		addresses, err := model.ReadAddress()
		if err != nil {
			fmt.Println(err)
		} else {
			c.JSON(200, addresses)
			fmt.Printf("%T\n", addresses)
		}
	})
	// GET タスク一覧取得
	router.GET("/task", func(c *gin.Context) {
		addresses, err := model.ReadTask()
		if err != nil {
			fmt.Println(err)
		} else {
			c.JSON(200, addresses)
			fmt.Printf("%T\n", addresses)
		}
	})

	//デバイス名に関してはタスクを追加するときについでに一緒に追加する

	// POST ステータスの追加
	router.POST("/status", func(c *gin.Context) {
		statusPatam := StatusParam{}
		c.BindJSON(&statusPatam)
		_, err := model.CreateStatus(statusPatam.StatusName)

		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(201, nil)
		}

	})

	// POST IPアドレスの追加
	router.POST("/address", func(c *gin.Context) {
		addressParam := AddressParam{}
		c.BindJSON(&addressParam)
		if addressParam.IPAddress == "" {
			c.JSON(201, nil)
			return
		}
		_, err := model.CreateAddress(addressParam.IPAddress)
		fmt.Println(err)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(201, nil)
		}

	})
	// POST IPアドレスの追加
	router.POST("/task", func(c *gin.Context) {
		taskParam := TaskParam{}
		c.BindJSON(&taskParam)
		_, err := model.CreateAddress(taskParam.IPAddress)
		fmt.Println(err)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(201, nil)
		}

	})

	//タスク一覧を受け取る

	// //タスクを追加する
	// router.POST("task",func(c *gin.Context)){

	// }

	// //タスクの内容を更新する
	// router.PUT("task/:id",func(c *gin.Context)){

	// }

	// //タスクを削除する
	// router.DELETE("task/:id",func(c *gin.Context)){

	// }

	return router
}

func main() {
	router := CreateRouter()
	router.Run(":" + os.Getenv("GO_PORT"))
}
