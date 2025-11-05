package models

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	Id          int64     `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"not null"`
	Content     string    `gorm:"not null"`
	UserID      uint      `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
	User        User
	IsPublished bool `gorm:"default:false"`
}

func (p *Post) BeforeCreate(tx *gorm.DB) error {
	return nil
}

// BeforeSave 文章保存前的钩子
func (p *Post) BeforeSave(tx *gorm.DB) error {
	// 自动生成摘要

	// 如果发布状态改变，更新发布时间
	return nil
}
