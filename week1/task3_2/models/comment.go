package models

// 评论
type Comment struct {
	Id         uint   `gorm:"primaryKey" json:"id"`
	Content    string `gorm:"size:100;not null" json:"content"`
	CommentNum int    `gorm:"size:100;not null" json:"commentNum"`
	PostId     uint   `gorm:"size:100;not null" json:"postId"`
}

func (c *Comment) TableName() string {
	return "comments"
}
