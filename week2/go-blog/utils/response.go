package utils

import "github.com/gin-gonic/gin"

// Response 统一响应格式
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// SuccessResponse 成功响应
func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse 错误响应
func ErrorResponse(c *gin.Context, statusCode int, errorMessage string) {
	c.JSON(statusCode, Response{
		Success: false,
		Error:   errorMessage,
	})
}
