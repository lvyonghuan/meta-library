package tool

import (
	"github.com/gin-gonic/gin"
	"log"
	"meta_library/util"
	"net/http"
	"time"
)

func GenerateGithubCookieAndSession(c *gin.Context) (sessionID string) {
	sessionID = RandString(10)
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("githubSessionID", sessionID, 3600, "/", "localhost", false, true) // 保存登陆状态
	return sessionID
}

func CookieWaiter(c *gin.Context) (isTimeout bool, err error, sessionID string) {
	log.Println("cookie计时器激活")
	timeout := 300 * time.Second       // 超时时间
	interval := 500 * time.Millisecond // 检查间隔
	start := time.Now()
	for time.Since(start) < timeout {
		if cookie, err := c.Request.Cookie("githubSessionID"); err == nil {
			sessionID = cookie.Value
			break
		}
		if err != nil {
			log.Println(err)
			util.NormErr(c, 99999, "我不道啊")
		}
		time.Sleep(interval)
	}
	if sessionID == "" {
		// cookie 未出现，超时
		util.NormErr(c, 60105, "等待超时")
		return true, err, ""
	}
	return false, err, sessionID
}
