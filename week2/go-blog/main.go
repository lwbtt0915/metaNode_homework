package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"web3/week2/go-blog/database"
	"web3/week2/go-blog/routes"
)

func main() {
	gin.SetMode(gin.DebugMode)

	database.InitDB()
	database.CreateAdminUser()
	r := routes.SetupRouter()

	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] | %3d | %13v | %15s | %-7s %s\n",
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Method,
			param.Path,
		)
	}))

	log.Println("博客后台服务启动")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("启动服务失败:", err)
	}
}
