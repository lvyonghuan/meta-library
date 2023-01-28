package model

type DiscussInfo struct {
	DiscussID int    `json:"discuss_id"`
	PostID    int    `json:"post_id"`
	ReplayID  int    `json:"replay_id"`
	Comment   string `json:"comment"`
	UserID    int    `json:"user_id"`
	PraiseNum int    `json:"praise_count"`
}
