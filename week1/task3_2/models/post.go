package models

// 文章
type Post struct {
	AuthorID uint      `gorm:"primaryKey" json:"authorId"`
	Title    string    `gorm:"size:100;not null" json:"title"`
	Comments []Comment `gorm:"foreignKey:CommentId" json:"comments"`
}

func (p *Post) TableName() string {
	return "post"
}
