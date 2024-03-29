package api

import (
	"github.com/gin-gonic/gin"
	"meta_library/service"
)

func InitRouter() {
	r := gin.Default()
	r.GET("/github_login", GithubLogin)
	user := r.Group("/user")
	{
		user.POST("/register", Register)
		user.GET("/token", Login)
		user.GET("/token/refresh", RefreshToken)
		user.PUT("/password", ChangePassword)
		user.GET("/info/:user_id", GetUserInfo)
		user.PUT("/info", ChangeUserInfo)
		user.GET("/redirect", service.RedirectGithub)
		user.GET("/link_github", LinkWithGithub)
		user.GET("/login_by_github", LoginByGithub)
	}
	book := r.Group("/book")
	{
		book.GET("/list", GetBookList)
		book.GET("/search", SearchBookInfo)
		book.PUT("/star", StarBook)
		book.GET("/label", GetSameLabelBook)
	}
	comment := r.Group("/comment")
	{
		comment.GET("/:book_id", GetCommentList)
		comment.POST("/:book_id", CreatComment)
		comment.DELETE("/:comment_id", DeleteComment)
		comment.PUT("/:comment_id", RefreshComment)
	}
	discuss := r.Group("/discuss")
	{
		discuss.POST("/:post_id", CreateDiscuss)
		discuss.GET("/:post_id", GetDiscussList)
		discuss.DELETE("/:discuss_id", DeleteDiscuss)
		discuss.POST("/replay/:discuss_id", ReplayDiscuss)
		discuss.GET("/check", CheckReplay)
		discuss.POST("/subscribe", SubscribeComment)
	}
	operate := r.Group("/operate")
	{
		operate.PUT("/praise/:target_id/model", Praise)
		operate.GET("/collect/list", GetCollectList)
		operate.PUT("/focus/:user_id", FocusUser)
	}
	r.Run()
}
