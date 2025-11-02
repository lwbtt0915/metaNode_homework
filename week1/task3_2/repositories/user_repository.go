package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"web3/week1/task3_2/models"
)

type BlogService struct {
	db *gorm.DB
}

type PostWithCommentStats struct {
	AuthorID uint   `json:"authorId"`
	Title    string `json:"title"`
	CId      string `json:"commentId"`
	Content  string `json:"content"`
}

// 查询某个用户发布的所有文章以及评论信息
func (blog *BlogService) GetUserPostWithComments(userId uint) ([]PostWithCommentStats, error) {
	var result []PostWithCommentStats

	query := "select *  from post p left join comments ct on ct.id=p.CommentId where p.authorId=?"
	err := blog.db.Raw(query, userId).Scan(&result).Error

	if err != nil {
		return nil, fmt.Errorf("获取用户发布文章和评论信息失败: %w", err)
	}

	return result, nil
}

// 评论数量最多的文章信息
func (blog *BlogService) GetMostCommentsPost() ([]models.Post, error) {
	var post []models.Post

	query := "SELECT\n\tp.*,\n\tCOUNT( c.id ) AS comment_count\nFROM\n\tposts p\n\tLEFT JOIN users u ON p.author_id = u.id \n\tLEFT JOIN comments c ON p.id = c.post_id \nWHERE\nGROUP BY\n\tp.id\nORDER BY\n\tcomment_count DESC"

	err := blog.db.Raw(query).Scan(&post).Error

	if err != nil {
		return nil, fmt.Errorf("获取评论数量最多的文章信息失败: %w", err)
	}

	return post, nil
}
