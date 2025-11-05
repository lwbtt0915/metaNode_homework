package controllers

import (
	"net/http"
	"web3/week2/go-blog/database"
	"web3/week2/go-blog/dto"
	"web3/week2/go-blog/models"
	"web3/week2/go-blog/utils"

	"github.com/gin-gonic/gin"
)

// GetProfile 获取用户资料
func GetProfile(c *gin.Context) {
	userID, _ := c.Get("userID")

	var user models.User
	result := database.DB.First(&user, userID)
	if result.Error != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}

	userResponse := dto.UserResponse{
		ID:       uint(user.Id),
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}

	utils.SuccessResponse(c, http.StatusOK, "Profile retrieved successfully", userResponse)
}

// UpdateProfile 更新用户资料
func UpdateProfile(c *gin.Context) {
	userID, _ := c.Get("userID")

	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Avatar   string `json:"avatar"`
		Bio      string `json:"bio"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request data")
		return
	}

	var user models.User
	result := database.DB.First(&user, userID)
	if result.Error != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}

	// 更新字段
	updates := make(map[string]interface{})
	if req.Username != "" && req.Username != user.Username {
		// 检查用户名是否已被使用
		var existingUser models.User
		if err := database.DB.Where("username = ? AND id != ?", req.Username, userID).First(&existingUser).Error; err == nil {
			utils.ErrorResponse(c, http.StatusConflict, "Username already exists")
			return
		}
		updates["username"] = req.Username
	}

	if req.Email != "" && req.Email != user.Email {
		// 检查邮箱是否已被使用
		var existingUser models.User
		if err := database.DB.Where("email = ? AND id != ?", req.Email, userID).First(&existingUser).Error; err == nil {
			utils.ErrorResponse(c, http.StatusConflict, "Email already exists")
			return
		}
		updates["email"] = req.Email
	}

	if req.Avatar != "" {
		updates["avatar"] = req.Avatar
	}

	if req.Bio != "" {
		updates["bio"] = req.Bio
	}

	if len(updates) > 0 {
		result = database.DB.Model(&user).Updates(updates)
		if result.Error != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update profile")
			return
		}
	}

	utils.SuccessResponse(c, http.StatusOK, "Profile updated successfully", nil)
}

// GetAllUsers 获取所有用户（管理员）
func GetAllUsers(c *gin.Context) {
	var users []models.User
	result := database.DB.Find(&users)
	if result.Error != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch users")
		return
	}

	var userResponses []dto.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, dto.UserResponse{
			ID:       uint(user.Id),
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
		})
	}

	utils.SuccessResponse(c, http.StatusOK, "Users retrieved successfully", userResponses)
}
