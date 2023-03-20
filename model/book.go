package model

type BookInfo struct {
	BookId      int    `json:"book_id" gorm:"primaryKey" gorm:"column:book_id"`
	Name        string `json:"name" gorm:"column:name"`
	IsStar      bool   `json:"is_star" gorm:"column:is_star"`
	Author      string `json:"author" gorm:"column:author"`
	CommentNum  string `json:"comment_num" gorm:"column:comment_num"`
	Score       int    `json:"score" gorm:"column:score"`
	Cover       string `json:"cover" gorm:"column:cover"`
	PublishTime string `json:"publish_time" gorm:"column:publish_time"`
	Link        string `json:"link" gorm:"column:link"`
	Label       string `json:"label" gorm:"column:label"`
}

type UserStar struct {
	UserID int  `json:"user_id" gorm:"primaryKey" gorm:"column:Id"`
	BookID int  `json:"book_id" gorm:"column:bookId"`
	IsStar bool `json:"is_star" gorm:"column:is_star"`
}
