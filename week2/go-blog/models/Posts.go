package models

type Post struct {
	Id      int64  `gorm:"primaryKey" json:"id"`
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
	UserID  uint
	User    User
}
