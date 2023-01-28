package api

import "github.com/gin-gonic/gin"

func InitRouter() {
	r := gin.Default()
	use := r.Group("/user")
	{
		use.POST("/register", Register)
		use.GET("/token", Login)
		use.GET("/token/refresh", RefreshToken)
		use.PUT("/password", ChangePassword)
		use.GET("/info/:user_id", GetUserInfo)
		use.PUT("/info", ChangeUserInfo)
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
	}
	operate := r.Group("/operate")
	{
		operate.PUT("/praise/:target_id/model", Praise)
		operate.GET("/collect/list", GetCollectList)
		operate.PUT("/focus/:user_id", FocusUser)
	}
	r.Run()
}
