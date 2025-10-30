package model

type Comment struct {
	Id      uint   `gorm:"primaryKey" json:"id"`
	Content string `gorm:"size:100;not null" json:"content"`
}
