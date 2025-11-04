package models

type Comments struct {
	Id      int64  `gorm:"primaryKey" json:"id"`
	Content string `gorm:"not null"`
	UserID  uint
	User    User
	PostID  uint
	Post    Post
}
