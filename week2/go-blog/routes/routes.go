package routes

import (
	"github.com/gin-gonic/gin"
	"web3/week2/go-blog/controllers"
	"web3/week2/go-blog/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 全局中间件
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.LoggerMiddleware())

	// 公共路由
	public := r.Group("/api")
	{
		auth := public.Group("/auth")
		{
			auth.POST("/login", controllers.Login)
			auth.POST("/register", controllers.Register)
		}
	}

	// 需要认证的路由
	authorized := r.Group("/api")
	authorized.Use(middleware.AuthMiddleware())
	{
		// 用户相关
		user := authorized.Group("/user")
		{
			user.GET("/profile", controllers.GetProfile)
			user.PUT("/profile", controllers.UpdateProfile)
		}

		// 文章管理
		posts := authorized.Group("/posts")
		{
			posts.POST("", controllers.CreatePost)
			posts.PUT("/:id", controllers.UpdatePost)
			posts.GET("/my-posts", controllers.GetMyPosts)
			posts.GET("/posts", controllers.GetPosts)
			posts.GET("/posts/:id", controllers.GetPost)
			posts.DELETE("/:id", controllers.DeletePost)
		}

		// 评论管理
		comments := authorized.Group("/comments")
		{
			comments.POST("created", controllers.CreateComment)
			comments.GET("/my-comments", controllers.GetMyComments)
			comments.GET("/byPostId", controllers.GetCommentsByPostID)
		}

		// 管理员路由
		admin := authorized.Group("/admin")
		admin.Use(middleware.AdminMiddleware())
		{
			admin.GET("/users", controllers.GetAllUsers)
		}
	}

	return r
}
