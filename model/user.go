package model

type UserInfo struct {
	Id              int    `gorm:"primaryKey" json:"id" gorm:"column:id"`
	UserName        string `json:"userName" gorm:"column:username"`
	PassWord        string `json:"passWord" gorm:"column:password"`
	Nickname        string `json:"nickname" gorm:"column:nickname"`
	Gender          string `json:"gender" gorm:"column:gender"`
	QQ              int    `json:"qq" gorm:"column:qq"`
	Birthday        string `json:"birthday" gorm:"column:birthday"`
	Email           string `json:"email" gorm:"column:email"`
	Avatar          string `json:"avatar" gorm:"column:avatar"`
	Introduction    string `json:"introduction" gorm:"column:introduction"`
	Phone           int    `json:"phone" gorm:"column:phone"`
	IsAdministrator bool   `json:"is_administrator" gorm:"is_administrator"`
}

type Follow struct {
	FollowerID int `gorm:"column:follower_id"`
	FolloweeID int `gorm:"column:followee_id"`
}
