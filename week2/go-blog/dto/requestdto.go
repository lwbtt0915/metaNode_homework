package dto

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterRequest 注册请求
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

// CreatePostRequest 创建文章请求
type CreatePostRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// UpdatePostRequest 更新文章请求
type UpdatePostRequest struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	IsPublished bool
}

// CreateCommentRequest 创建评论请求
type CreateCommentRequest struct {
	Content  string `json:"content" binding:"required"`
	PostID   uint   `json:"postId" binding:"required"`
	ParentId uint
}

type GetCommentsRequest struct {
	PostID uint `json:"postId" binding:"required"`
}

// PaginationRequest 分页请求
type PaginationRequest struct {
	Page     int `form:"page" binding:"min=1"`
	PageSize int `form:"page_size" binding:"min=1,max=100"`
}
