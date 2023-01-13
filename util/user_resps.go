package util

import (
	"github.com/gin-gonic/gin"
	"meta_library/model"
	"net/http"
)

type TokenResponse struct {
	Status int       `json:"status"`
	Info   string    `json:"info"`
	Data   TokenData `json:"data"`
}

type TokenData struct {
	RefreshToken string `json:"refresh_token"`
	Token        string `json:"token"`
}

type userInfo struct {
	ID           int    `json:"id"`
	Avatar       string `json:"avatar"`
	Nickname     string `json:"nickname"`
	Introduction string `json:"introduction"`
	Phone        int    `json:"phone"`
	QQ           int    `json:"qq"`
	Gender       string `json:"gender"`
	Email        string `json:"email"`
	Birthday     string `json:"birthday"`
}

type Data struct {
	User userInfo `json:"user"`
}

type userInfoResponse struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
	Data   Data   `json:"data"`
}

func RespSuccess(c *gin.Context, token string, refreshToken string) { //token获取成功
	response := TokenResponse{
		Status: 10000,
		Info:   "success",
		Data: TokenData{
			RefreshToken: refreshToken,
			Token:        token,
		},
	}

	c.JSON(http.StatusOK, response)
}

func RespUserInfoSuccess(c *gin.Context, u model.UserInfo) { //用户信息获取成功
	response := userInfoResponse{
		Status: 10000,
		Info:   "success",
		Data: Data{User: userInfo{
			ID:           u.Id,
			Avatar:       u.Avatar,
			Nickname:     u.Nickname,
			Introduction: u.Introduction,
			Phone:        u.Phone,
			QQ:           u.QQ,
			Gender:       u.Gender,
			Email:        u.Email,
			Birthday:     u.Birthday,
		}},
	}
	c.JSON(http.StatusOK, response)
}
