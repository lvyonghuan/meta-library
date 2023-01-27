package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"meta_library/model"
	"meta_library/service"
	"meta_library/tool"
	"meta_library/util"
	"strconv"
)

func CreateDiscuss(c *gin.Context) {
	token := c.GetHeader("Authorization")
	postIDString := c.Param("post_id")
	comment := c.PostForm("comment")
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
	uUser, err := service.SearchUserByUserName(username)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	postID, err := strconv.Atoi(postIDString)
	if err != nil {
		log.Printf("search discuss error:%v", err)
		util.NormErr(c, 70012, "post_id非法")
		return
	}
	var uDiscuss model.DiscussInfo
	uDiscuss.PostID = postID
	uDiscuss.Comment = comment
	uDiscuss.UserID = uUser.Id
	discussID, err := service.CreateDiscuss(uDiscuss)
	if err != nil {
		log.Printf("search discuss error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.CreatDiscussRespSuccess(c, discussID)
}
