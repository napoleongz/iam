package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建一个默认的Gin引擎
	r := gin.Default()

	// 定义一个GET请求的路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, welcome to the Gin demo!",
		})
	})

	// 启动HTTP服务器，监听在本地的8080端口
	r.Run(":8080")
}

func Testa() {
	fmt.Println("test")

}

type Test struct {
	Name string
	Age  int
}

func Testb(t Test) {
	fmt.Println(t)
}
