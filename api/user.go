package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"meta_library/model"
	"meta_library/service"
	"meta_library/tool"
	"meta_library/util"
	"net/http"
	"strconv"
	"time"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		util.RespParamErr(c)
		return
	}
	//到30行：用户名查重
	u, err := service.SearchUserByUserName(username)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	if u.UserName != "" {
		util.NormErr(c, 60001, "用户名已注册")
		return
	}
	err = service.InsertUser(model.UserInfo{
		UserName: username,
		PassWord: password,
	})
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.RespOK(c)
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	u, err := service.SearchUserByUserName(username) //查找用户名是否存在于数据库中,且把密码从数据库取出来
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	if u.UserName == "" {
		util.NormErr(c, 60002, "用户未注册或用户名输入错误")
		return
	}
	if u.PassWord != password {
		util.NormErr(c, 60003, "密码错误")
		return
	}
	token, err := tool.GenerateToken([]byte("114"), 3600*time.Second, username)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 60010, "登陆错误")
		return
	}
	refreshToken, err := tool.GenerateToken([]byte("514"), 86400*time.Second, username) //114和514都是签名用的秘钥，114是token的，514是refresh_token的
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 60010, "登陆错误")
		return
	}
	util.RespSuccess(c, token, refreshToken)
}

func RefreshToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	refreshToken := c.Query("refresh_token")
	_, username, err := tool.TokenExpired([]byte("114"), token)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 60100, "token错误")
		return
	}
	isTure, _, err := tool.TokenExpired([]byte("514"), refreshToken)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 60100, "token错误")
		return
	} else if err == nil && isTure == false {
		util.NormErr(c, 60103, "refresh_token过期，请重新登陆")
		return
	}
	token, err = tool.GenerateToken([]byte("114"), 3600*time.Second, username)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 600104, "token申请错误")
		return
	}
	refreshToken, err = tool.GenerateToken([]byte("514"), 86400*time.Second, username)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 60104, "token申请错误")
		return
	}
	util.RespSuccess(c, token, refreshToken)
}

func ChangePassword(c *gin.Context) {
	//获取用户信息
	token := c.GetHeader("Authorization")
	oldPassword := c.Query("old_password")
	newPassword := c.Query("new_password")
	//解密token
	isExpired, username, err := tool.TokenExpired([]byte("114"), token)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 600100, "token错误")
		return
	}
	if !isExpired {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 600102, "token已过期")
		return
	}
	//查找用户信息
	u, err := service.SearchUserByUserName(username)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	if u.PassWord != oldPassword {
		util.NormErr(c, 60003, "密码错误")
		return
	}
	err = service.ChangePasswordByUsername(username, newPassword)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.RespOK(c)
}

func GetUserInfo(c *gin.Context) {
	idString := c.Param("user_id")
	fmt.Println(idString)
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 60004, "UID非法")
		return
	}
	u, err := service.SearchUserByUserId(id)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.RespUserInfoSuccess(c, u)
}

func ChangeUserInfo(c *gin.Context) {
	token := c.GetHeader("Authorization")
	nickname := c.PostForm("nickname")
	avatar := c.PostForm("avatar")
	introduction := c.PostForm("introduction")
	telephoneString := c.PostForm("telephone")
	telephone, err := strconv.Atoi(telephoneString)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 60011, "非法数值")
		return
	}
	qqString := c.PostForm("qq")
	qq, err := strconv.Atoi(qqString)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 60011, "非法数值")
		return
	}
	gender := c.PostForm("gender")
	email := c.PostForm("email")
	birthday := c.PostForm("birthday")
	isExist, username, err := tool.TokenExpired([]byte("114"), token)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 600100, "token错误")
		return
	}
	if !isExist {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 600102, "token已过期")
		return
	}
	u, err := service.SearchUserByUserName(username)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	u.Nickname = nickname
	u.Avatar = avatar
	u.Introduction = introduction
	u.Phone = telephone
	u.QQ = qq
	u.Gender = gender
	u.Email = email
	u.Birthday = birthday
	err = service.ChangeUserInfo(u)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.RespOK(c)
}

func RedirectGithub(c *gin.Context) {
	// 构造授权 URL
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

func GithubLogin(c *gin.Context) {
	// 从查询参数中获取授权代码和状态令牌
	code := c.Query("code")
	//state := c.Query("state")
	// 根据授权代码交换访问令牌
	token, err := service.GetAccessToken(code)
	if err != nil {
		// 处理错误
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	// 使用访问令牌获取用户数据
	user, err := service.GetUserData(token)
	if err != nil {
		// 处理错误
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	// 响应成功
	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"user":    user,
	})
}
