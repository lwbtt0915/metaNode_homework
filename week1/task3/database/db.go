package database

import (
	"dataDriver/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

var DB *sqlx.DB

// InitDB 初始化数据库连接
func InitDB() error {
	cfg := config.LoadDataBaseConfig()
	connStr := cfg.GetConnectionsString()

	db, err := sqlx.Connect("mysql", connStr)
	if err != nil {
		return fmt.Errorf("无法连接数据库: %w", err)
	}

	// 配置连接池
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * 60 * 60) // 5小时

	// 测试连接
	if err := db.Ping(); err != nil {
		return fmt.Errorf("数据库连接测试失败: %w", err)
	}

	DB = db
	log.Println("数据库连接成功")
	return nil
}

// closeDB 关闭数据库连接
func CloseDB() error {
	if DB != nil {
		return DB.Close()
	}

	return nil
}

// GetDB 数据库实例
func GetDB() *sqlx.DB {
	return DB
}
