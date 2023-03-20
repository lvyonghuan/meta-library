package model

type Conf struct {
	ClientId     string
	ClientSecret string
	RedirectUrl  string
	State        string
}

type GithubRelate struct {
	GithubID int `gorm:"column:github_id"`
	UID      int `gorm:"column:uid"`
}

type Session struct {
	Session string `gorm:"column:session"`
	Token   string `gorm:"column:token"`
}
