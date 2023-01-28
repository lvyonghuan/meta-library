package model

type UserInfo struct {
	Id              int    `json:"id"`
	UserName        string `json:"userName"`
	PassWord        string `json:"passWord"`
	Nickname        string `json:"nickname"`
	Gender          string `json:"gender"`
	QQ              int    `json:"qq"`
	Birthday        string `json:"birthday"`
	Email           string `json:"email"`
	Avatar          string `json:"avatar"`
	Introduction    string `json:"introduction"`
	Phone           int    `json:"phone"`
	IsAdministrator bool   `json:"is_administrator"`
}
