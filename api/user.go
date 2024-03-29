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
	err = service.ChangePasswordByUsername(u.Id, newPassword)
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

func GithubLogin(c *gin.Context) {
	log.Println("GithubLogin function called")
	// 从查询参数中获取授权代码和状态令牌
	code := c.Query("code")
	//state := c.Query("state")
	// 根据授权代码交换访问令牌
	token, err := service.GetAccessToken(code)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("令牌获取")
	sessionID := tool.GenerateGithubCookieAndSession(c)
	log.Println("cookie设置完成")
	err = service.StoreSession(sessionID, token)
	if err != nil {
		log.Println(err)
		util.RsepInternalErr(c)
		return
	}
	err = tool.DeleteSessionTimer(sessionID)
	if err != nil {
		util.RsepInternalErr(c)
		return
	}
	log.Println("github设置模块运行完毕")
}

func LinkWithGithub(c *gin.Context) {
	token := c.GetHeader("Authorization")
	u, err := service.CheckToken(token, c)
	if err != nil {
		return
	}
	service.RedirectGithub(c)
	isTimeout, err, sessionID := tool.CookieWaiter(c)
	if isTimeout || err != nil {
		return
	}
	token, err = service.SearchSessionByID(sessionID)
	if err != nil {
		log.Println(err)
		util.RsepInternalErr(c)
		return
	}
	user, err := service.GetUserData(token)
	if err != nil {
		// 处理错误
		log.Println(err)
		util.NormErr(c, 99999, "我不道啊")
		return
	}
	userID := user["id"].(float64)
	err = service.LinkWithGithub(int(userID), u.Id)
	if err != nil {
		log.Println(err)
		util.RsepInternalErr(c)
		return
	}
	util.RespOK(c)
}

func LoginByGithub(c *gin.Context) {
	service.RedirectGithub(c)
	isTimeout, err, sessionID := tool.CookieWaiter(c)
	if isTimeout || err != nil {
		return
	}
	log.Println("我跑这里来了")
	githubToken, err := service.SearchSessionByID(sessionID)
	if err != nil {
		log.Println(err)
		util.RsepInternalErr(c)
		return
	}
	user, err := service.GetUserData(githubToken)
	if err != nil {
		// 处理错误
		log.Println(err)
		util.NormErr(c, 99999, "我不道啊")
		return
	}
	userID := user["id"].(float64)
	uid, err := service.LoginByGithub(int(userID))
	u, err := service.SearchUserByUserId(uid)
	if err != nil {
		log.Println(err)
		util.RsepInternalErr(c)
		return
	}
	log.Println(u)
	token, err := tool.GenerateToken([]byte("114"), 3600*time.Second, u.UserName)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 60010, "登陆错误")
		return
	}
	refreshToken, err := tool.GenerateToken([]byte("514"), 86400*time.Second, u.UserName)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 60010, "登陆错误")
		return
	}
	util.RespSuccess(c, token, refreshToken)
}
