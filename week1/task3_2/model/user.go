package model

type User struct {
	id    int64  `gorm:"primaryKey" json:"id"`
	name  string `gorm:"size:100;not null" json:"name"`
	Posts []Post `gorm:"foreignKey:AuthorID" json:"posts"`
}
