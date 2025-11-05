package main

import (
	"log"
	"web3/week2/go-blog/database"
	"web3/week2/go-blog/routes"
)

func main() {
	database.InitDB()
	database.CreateAdminUser()
	r := routes.SetupRouter()

	log.Println("博客后台服务启动")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("启动服务失败:", err)
	}
}
