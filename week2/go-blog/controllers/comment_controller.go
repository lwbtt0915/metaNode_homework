package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web3/week2/go-blog/database"
	"web3/week2/go-blog/dto"
	"web3/week2/go-blog/models"
	"web3/week2/go-blog/utils"
)

// CreateComment 创建评论
func CreateComment(c *gin.Context) {
	userID, _ := c.Get("userID")

	var req dto.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 检查文章是否存在
	var post models.Post
	result := database.DB.First(&post, req.PostID)
	if result.Error != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Post not found")
		return
	}

	comment := models.Comments{
		Content:  req.Content,
		UserId:   userID.(uint),
		PostId:   req.PostID,
		ParentID: req.ParentId,
	}

	result = database.DB.Create(&comment)
	if result.Error != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create comment")
		return
	}

	// 重新加载关联数据
	database.DB.Preload("User").Preload("Post").First(&comment, comment.Id)

	utils.SuccessResponse(c, http.StatusCreated, "Comment created successfully", comment)
}

// GetMyComments 获取我的评论
func GetMyComments(c *gin.Context) {
	userID, _ := c.Get("userID")

	var comments []models.Comments
	result := database.DB.Where("user_id = ?", userID).
		Preload("Post").
		Order("created_at DESC").
		Find(&comments)

	if result.Error != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch comments")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Comments retrieved successfully", comments)
}

// 根据文章ID查询评论
func GetCommentsByPostID(c *gin.Context) {
	var req dto.GetCommentsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var comments []models.Comments
	result := database.DB.Where("post_id = ?", req.PostID).Preload("Post").Order("created_at DESC").
		Find(&comments)

	if result.Error != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch comments")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Comments query successfully", comments)
}
