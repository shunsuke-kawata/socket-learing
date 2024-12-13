package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
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
		AllowCredentials: true,
		AllowWebSockets:  true,
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
		fmt.Println(statuses)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, nil)
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
			c.JSON(500, nil)
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
			c.JSON(500, nil)
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
	// POST IPアドレスの追加
	router.POST("/address", func(c *gin.Context) {
		addressParam := AddressParam{}
		c.BindJSON(&addressParam)

		newAddress, err := model.CreateAddress(addressParam.IPAddress)
		if err != nil {
			if err.Error() == "address already exists" {
				c.JSON(400, nil) // すでに存在する場合
			} else {
				c.JSON(500, err.Error()) // その他のエラー
			}
		} else {
			c.JSON(201, newAddress) // 新しく作成したアドレスを返す
		}
	})

	// POST IPアドレスの追加
	// POST タスクの追加
	router.POST("/task", func(c *gin.Context) {
		taskParam := TaskParam{}
		c.BindJSON(&taskParam)

		//IPアドレスをヘッダーから取得
		clientIP := c.GetHeader("X-Forwarded-For")
		if clientIP == "" {
			clientIP = c.ClientIP()
		} else {
			fmt.Println("forword", clientIP)
		}
		newTask, err := model.CreateTask(taskParam.Title, taskParam.Description, clientIP)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(201, newTask)
		}
	})

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

	//socket通信
	router.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println("WebSocket upgrade error:", err)
			c.String(http.StatusBadRequest, "Failed to upgrade connection")
			return
		}
		fmt.Println("WebSocket connection established.")
		defer conn.Close()

		for {
			messageType, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Error reading message:", err)
				break
			}
			fmt.Printf("Received message: %s\n", msg)
			err = conn.WriteMessage(messageType, msg)
			if err != nil {
				fmt.Println("Error writing message:", err)
				break
			}
		}
	})

	router.Run(":" + os.Getenv("GO_PORT"))
}
