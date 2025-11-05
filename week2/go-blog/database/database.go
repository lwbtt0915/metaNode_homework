package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"web3/week2/go-blog/models"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() {
	dsn := "root:12345678@tcp(localhost:3306)/web3?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}

	// 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("获取数据库连接失败:", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	DB = db

	// 自动迁移
	autoMigrate()

	log.Println("✅ 数据库连接成功!")
}

// autoMigrate 自动迁移表结构
func autoMigrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Comments{},
	)

	if err != nil {
		log.Fatal("自动迁移失败:", err)
	}

	log.Println("✅ 表结构迁移成功!")
}

// CreateAdminUser 创建管理员用户（可选）
func CreateAdminUser() {
	var count int64
	DB.Model(&models.User{}).Where("role = ?", "admin").Count(&count)

	if count == 0 {
		admin := models.User{
			Username: "admin",
			Email:    "admin@blog.com",
			Role:     "admin",
		}

		err := admin.HashPassword("admin123")
		if err != nil {
			log.Fatal("加密密码失败:", err)
		}

		result := DB.Create(&admin)
		if result.Error != nil {
			log.Fatal("创建管理员失败:", result.Error)
		}

		log.Println("✅ 默认管理员创建成功!")
	}
}
