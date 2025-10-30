package model

type Post struct {
	AuthorID uint      `gorm:"primaryKey" json:"authorId"`
	Title    string    `gorm:"size:100;not null" json:"title"`
	Comments []Comment `gorm:"foreignKey:CommentId" json:"comments"`
}
