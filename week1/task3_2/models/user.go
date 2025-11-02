package models

type User struct {
	Id    int64  `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"size:100;not null" json:"name"`
	Posts []Post `gorm:"foreignKey:AuthorID" json:"posts"`
}

func (u *User) TableName() string {
	return "user"
}
