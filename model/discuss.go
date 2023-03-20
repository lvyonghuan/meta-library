package model

type DiscussInfo struct {
	DiscussID int    `json:"discuss_id" gorm:"primaryKey" gorm:"column:discuss_id"`
	PostID    int    `json:"post_id" gorm:"column:post_id"`
	ReplayID  int    `json:"replay_id" gorm:"column:replay_id"`
	Comment   string `json:"comment" gorm:"column:comment"`
	UserID    int    `json:"user_id" gorm:"column:user_id"`
	PraiseNum int    `json:"praise_count" gorm:"column:praise_count"`
	ReplayUid int    `json:"replay_uid" gorm:"column:replay_uid"`
}
