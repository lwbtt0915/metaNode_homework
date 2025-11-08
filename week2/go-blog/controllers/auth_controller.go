package controllers

import (
	"net/http"
	"web3/week2/go-blog/database"
	"web3/week2/go-blog/dto"
	"web3/week2/go-blog/models"
	"web3/week2/go-blog/utils"

	"github.com/gin-gonic/gin"
)

// Login 用户登录
func Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request data")
		return
	}

	var user models.User
	result := database.DB.Where("username = ? OR email = ?", req.Username, req.Username).First(&user)
	if result.Error != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	// 检查密码
	if err := user.CheckPassword(req.Password); err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	// 检查用户是否激活
	if !user.IsActive {
		utils.ErrorResponse(c, http.StatusForbidden, "Account is deactivated")
		return
	}

	// 生成token
	token, err := utils.GenerateToken(uint(user.Id), user.Username, user.Role)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	userResponse := dto.UserResponse{
		ID:       uint(user.Id),
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}

	utils.SuccessResponse(c, http.StatusOK, "Login successful", gin.H{
		"user":  userResponse,
		"token": token,
	})
}

// Register 用户注册
func Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 检查用户名是否已存在
	var existingUser models.User
	result := database.DB.Where("username = ? OR email = ?", req.Username, req.Email).First(&existingUser)
	if result.Error == nil {
		utils.ErrorResponse(c, http.StatusConflict, "Username or email already exists")
		return
	}

	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Role:     "manager",
	}

	// 加密密码
	if err := user.HashPassword(req.Password); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	// 创建用户
	result = database.DB.Create(&user)
	if result.Error != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	userResponse := dto.UserResponse{
		ID:       uint(user.Id),
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}

	utils.SuccessResponse(c, http.StatusCreated, "User registered successfully", userResponse)
}
