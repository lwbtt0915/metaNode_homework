package models

import "time"

type Comments struct {
	Id        int64     `gorm:"primaryKey；autoIncrement" json:"id"`
	Content   string    `gorm:"not null" json:"content"`
	UserId    uint      `gorm:"not null" json:"userId"`
	PostId    uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	//User      User
	//Post     Post
	ParentID uint `gorm:"not null" json:"parentId"`
}
