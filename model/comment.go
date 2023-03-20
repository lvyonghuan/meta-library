package model

type CommentInfo struct {
	PostID      int    `json:"post_id" gorm:"primaryKey" gorm:"column:post_id"`
	PublishTime string `json:"publish_time" gorm:"column:book_id"`
	Content     string `json:"content" gorm:"column:publish_time"`
	UserID      int    `json:"user_id" gorm:"column:content"`
	Avatar      string `json:"avatar" gorm:"column:user_id"`
	Nickname    string `json:"nickname" gorm:"column:avatar"`
	PraiseCount int    `json:"praise_count" gorm:"column:nickname"`
	IsPraised   bool   `json:"is_praised" gorm:"column:praise_count"`
	IsFocus     bool   `json:"is_focus" gorm:"column:is_praised"`
	BookID      int    `json:"book_id" gorm:"column:is_focus"`
}
