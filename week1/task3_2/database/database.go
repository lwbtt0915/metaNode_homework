package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gormboke/model"
	"log"
)

var DB *gorm.DB

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func InitDB(cfg Config) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		return fmt.Errorf("无法连接到数据库： %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("获取数据库实例失败：%w", err)
	}

	//连接池配置
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * 60 * 60)

	DB = db
	log.Println("数据库连接成功")

	//自动迁移
	if err := AutoMigrate(); err != nil {
		return fmt.Errorf("数据库迁移失败： %w", err)
	}

	log.Println("数据库表创建成功")
	return nil
}

func AutoMigrate() error {
	return DB.AutoMigrate(
		&model.User{},
		&model.Post{},
		&model.Comment{},
	)
}

func GetDB() *gorm.DB {
	return DB
}

func CloseDB() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
