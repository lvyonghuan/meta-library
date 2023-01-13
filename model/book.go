package model

type BookInfo struct {
	BookId      int    `json:"book_id"`
	Name        string `json:"name"`
	IsStar      bool   `json:"is_star"`
	Author      string `json:"author"`
	CommentNum  string `json:"comment_num"`
	Score       int    `json:"score"`
	Cover       string `json:"cover"`
	PublishTime string `json:"publish_time"`
	Link        string `json:"link"`
	Label       string `json:"label"`
}

type UserStar struct {
	UserID int  `json:"user_id"`
	BookID int  `json:"book_id"`
	IsStar bool `json:"is_star"`
}
