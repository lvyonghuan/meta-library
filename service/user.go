package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"meta_library/dao"
	"meta_library/model"
	"meta_library/tool"
	"meta_library/util"
	"net/http"
	"net/url"
)

func SearchUserByUserName(username string) (u model.UserInfo, err error) {
	u, err = dao.SearchUserByUserName(username)
	return u, err
}

func InsertUser(u model.UserInfo) (err error) {
	err = dao.InsertUser(u)
	return err
}

func ChangePasswordByUsername(username string, newPassword string) (err error) {
	err = dao.ChangePasswordByUsername(username, newPassword)
	return err
}

func SearchUserByUserId(id int) (u model.UserInfo, err error) {
	u, err = dao.SearchUserByUserId(id)
	return u, err
}

func ChangeUserInfo(u model.UserInfo) (err error) {
	err = dao.ChangeUserInfo(u)
	return err
}

func GetAccessToken(code string) (string, error) {
	// 构造获取访问令牌的请求
	conf := model.Conf{
		ClientId:     "Iv1.993fdcaba2e1356f",
		ClientSecret: "d4fa07c0d67b6f8d8f9ee8341748949e2cde6ce4",
		RedirectUrl:  "http://localhost:8080/github_login",
	}
	data := url.Values{}
	data.Set("client_id", conf.ClientId)
	data.Set("client_secret", conf.ClientSecret)
	data.Set("code", code)
	data.Set("redirect_uri", conf.RedirectUrl)
	//data.Set("state", state)
	// 发送请求并解析响应
	resp, err := http.PostForm("https://github.com/login/oauth/access_token", data)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("github 获取 token 模块问题:", err)
			return
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	// 解析响应，提取访问令牌
	values, err := url.ParseQuery(string(body))
	if err != nil {
		return "", err
	}
	token := values.Get("access_token")
	if token == "" {
		return "", errors.New("access token not found")
	}
	return token, nil
}

func GetUserData(token string) (map[string]interface{}, error) {
	// 构造获取用户数据的请求
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	// 发送请求并解析响应
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("github 获取用户信息模块错误:", err)
		}
	}(resp.Body)
	// 读取响应体数据
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}
	// 解析响应，提取用户
	var user map[string]interface{}
	err = json.Unmarshal(buf.Bytes(), &user)
	if err != nil {
		return nil, err
	}
	return user, err
}

func CheckToken(token string, c *gin.Context) (u model.UserInfo, err error) {
	isExist, username, err := tool.TokenExpired([]byte("114"), token)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 600100, "token错误")
		return
	}
	if !isExist {
		util.NormErr(c, 600102, "token已过期")
		return
	}
	u, err = SearchUserByUserName(username)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	return u, err
}

func LinkWithGithub(githubID int, uID int) (err error) {
	err = dao.LinkWithGithub(githubID, uID)
	return
}

func RedirectGithub(c *gin.Context) {
	// 构造授权 URL
	log.Println("RedirectGithub function called")
	state := tool.GenerateState()
	conf := model.Conf{
		ClientId:     "Iv1.993fdcaba2e1356f",
		ClientSecret: "d4fa07c0d67b6f8d8f9ee8341748949e2cde6ce4",
		RedirectUrl:  "http://localhost:8080/github_login",
		State:        state,
	}
	//url := "https://github.com/login/oauth/authorize?client_id=" + conf.ClientId + "&redirect_uri=" + conf.RedirectUrl + "&state=" + state
	url := "https://github.com/login/oauth/authorize?client_id=" + conf.ClientId + "&redirect_uri=" + conf.RedirectUrl
	// 重定向到授权 URL
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func LoginByGithub(githubID int) (uid int, err error) {
	uid, err = dao.SearchGithubID(githubID)
	return uid, err
}

func StoreSession(sessionID string, value string) (err error) {
	err = dao.StoreSession(sessionID, value)
	return
}

func SearchSessionByID(sessionID string) (token string, err error) {
	token, err = dao.SearchSessionByID(sessionID)
	return token, err
}
