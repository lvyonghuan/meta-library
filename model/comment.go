package model

type CommentInfo struct {
	PostID      int    `json:"post_id"`
	PublishTime string `json:"publish_time"`
	Content     string `json:"content"`
	UserID      int    `json:"user_id"`
	Avatar      string `json:"avatar"`
	Nickname    string `json:"nickname"`
	PraiseCount int    `json:"praise_count"`
	IsPraised   bool   `json:"is_praised"`
	IsFocus     bool   `json:"is_focus"`
	BookID      int    `json:"book_id"`
}
