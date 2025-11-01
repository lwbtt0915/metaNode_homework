package main

import (
	"fmt"
	"gormboke/database"
	"gormboke/model"
	"log"
	"os"
	"sync"
	"time"
)

func main1() {
	m := make(map[string]int)
	var lock sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		for {
			lock.Lock()
			m["a"]++
			lock.Unlock()
		}
	}()

	go func() {
		for {
			lock.Lock()
			m["a"]++
			fmt.Println(m["a"])
			lock.Unlock()
		}
	}()

	select {
	case <-time.After(time.Second * 5):
		fmt.Println("timeout, stopping")
	}
}

func getEnv(key, defalutValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defalutValue
}

func main() {
	config := database.Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "3306"),
		User:     getEnv("DB_USER", "root"),
		Password: getEnv("DB_PASSWORD", "password"),
		DBName:   getEnv("DB_NAME", "blog_system"),
	}

	if err := database.InitDB(config); err != nil {
		log.Fatalf("数据库初始化失败：%v", err)
	}

	defer database.CloseDB()

	fmt.Println("表创建成功")

	if err := createSampleData(); err != nil {
		log.Printf("创建示例数据失败: %v", err)
	}

}

func createSampleData() error {
	db := database.GetDB()

	users := []model.User{{Id: 1, Name: "tim"}, {Id: 2, Name: "jeffry"}}

	if err := db.Create(&users).Error; err != nil {
		return fmt.Errorf("创建用户失败： %w", err)
	}

	posts := model.Post{
		AuthorID: uint(users[1].Id),
		Title:    "经济",
	}

	if err := db.Create(&posts).Error; err != nil {
		return fmt.Errorf("创建投递失败： %w", err)
	}
	return nil
}
