package controllers

import (
	"net/http"
	"strconv"
	"web3/week2/go-blog/database"
	"web3/week2/go-blog/dto"
	"web3/week2/go-blog/models"
	"web3/week2/go-blog/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetPosts 获取文章列表
func GetPosts(c *gin.Context) {
	var pagination dto.PaginationRequest
	if err := c.ShouldBindQuery(&pagination); err != nil {
		pagination.Page = 1
		pagination.PageSize = 10
	}

	if pagination.Page < 1 {
		pagination.Page = 1
	}
	if pagination.PageSize < 1 || pagination.PageSize > 100 {
		pagination.PageSize = 10
	}

	offset := (pagination.Page - 1) * pagination.PageSize

	var posts []models.Post
	var total int64

	// 只查询已发布的文章
	query := database.DB.Where("is_published = ?", true)

	// 获取总数
	query.Model(&models.Post{}).Count(&total)

	// 获取分页数据
	result := query.Offset(offset).Limit(pagination.PageSize).
		Order("created_at DESC").
		Find(&posts)

	if result.Error != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch posts")
		return
	}

	paginationResponse := dto.PaginationResponse{
		Page:      pagination.Page,
		PageSize:  pagination.PageSize,
		Total:     int(total),
		TotalPage: (int(total) + pagination.PageSize - 1) / pagination.PageSize,
	}

	utils.SuccessResponse(c, http.StatusOK, "Posts retrieved successfully", gin.H{
		"posts":      posts,
		"pagination": paginationResponse,
	})
}

// GetPost 获取单篇文章
func GetPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid post ID")
		return
	}

	var post models.Post
	result := database.DB.First(&post, id)

	if result.Error != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Post not found")
		return
	}

	// 增加阅读数
	if post.IsPublished {
		database.DB.Model(&post).Update("view_count", gorm.Expr("view_count + ?", 1))
	}

	utils.SuccessResponse(c, http.StatusOK, "Post retrieved successfully", post)
}

// CreatePost 创建文章
func CreatePost(c *gin.Context) {
	userID, _ := c.Get("userID")

	var req dto.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	post := models.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID.(uint),
	}

	result := database.DB.Create(&post)
	if result.Error != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create post")
		return
	}

	// 重新加载关联数据
	database.DB.Preload("User").Preload("Tags").Preload("Categories").First(&post, post.Id)

	utils.SuccessResponse(c, http.StatusCreated, "Post created successfully", post)
}

// UpdatePost 更新文章
func UpdatePost(c *gin.Context) {
	userID, _ := c.Get("userID")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid post ID")
		return
	}

	var post models.Post
	result := database.DB.Preload("Tags").Preload("Categories").First(&post, id)
	if result.Error != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Post not found")
		return
	}

	// 检查权限（只能更新自己的文章，或者管理员）
	role, _ := c.Get("role")
	if post.UserID != userID.(uint) && role != "admin" {
		utils.ErrorResponse(c, http.StatusForbidden, "No permission to update this post")
		return
	}

	var req dto.UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	updates := make(map[string]interface{})
	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Content != "" {
		updates["content"] = req.Content
	}

	updates["is_published"] = req.IsPublished

	if len(updates) > 0 {
		database.DB.Model(&post).Updates(updates)
	}

	// 重新加载数据
	database.DB.Preload("User").Preload("Tags").Preload("Categories").First(&post, post.Id)

	utils.SuccessResponse(c, http.StatusOK, "Post updated successfully", post)
}

// GetMyPosts 获取我的文章
func GetMyPosts(c *gin.Context) {
	userID, _ := c.Get("userID")

	var pagination dto.PaginationRequest
	if err := c.ShouldBindQuery(&pagination); err != nil {
		pagination.Page = 1
		pagination.PageSize = 10
	}

	offset := (pagination.Page - 1) * pagination.PageSize

	var posts []models.Post
	var total int64

	query := database.DB.Where("user_id = ?", userID).Preload("Tags").Preload("Categories")
	query.Model(&models.Post{}).Count(&total)

	result := query.Offset(offset).Limit(pagination.PageSize).
		Order("created_at DESC").
		Find(&posts)

	if result.Error != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch posts")
		return
	}

	paginationResponse := dto.PaginationResponse{
		Page:      pagination.Page,
		PageSize:  pagination.PageSize,
		Total:     int(total),
		TotalPage: (int(total) + pagination.PageSize - 1) / pagination.PageSize,
	}

	utils.SuccessResponse(c, http.StatusOK, "Posts retrieved successfully", gin.H{
		"posts":      posts,
		"pagination": paginationResponse,
	})
}
