package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"meta_library/dao/rds"
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
	rds.Publish(c, postID, comment)
	util.CreatDiscussRespSuccess(c, discussID)
}

func GetDiscussList(c *gin.Context) { //获取一个帖子下全部的回复信息
	postIDString := c.Param("post_id")
	postID, err := strconv.Atoi(postIDString)
	if err != nil {
		log.Printf("search discuss error:%v", err)
		util.NormErr(c, 70012, "post_id非法")
		return
	}
	uDiscussInfo, err := service.GetDiscussList(postID)
	if err != nil {
		log.Printf("search discuss error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.GetDiscussInfoSuccess(c, uDiscussInfo)
}

func DeleteDiscuss(c *gin.Context) {
	token := c.GetHeader("Authorization")
	discussIDString := c.Param("discuss_id")
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
	discussID, err := strconv.Atoi(discussIDString)
	if err != nil {
		log.Printf("search discuss error:%v", err)
		util.NormErr(c, 70013, "discuss_id非法")
		return
	}
	err = service.DeleteDiscuss(discussID, uUser.Id, uUser.IsAdministrator)
	if err != nil {
		if err.Error() == "discuss_id and user_id not match" {
			util.NormErr(c, 70014, "用户无权限删除此书评")
			return
		}
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.RespOK(c)
}

func ReplayDiscuss(c *gin.Context) {
	token := c.GetHeader("Authorization")
	discussIDString := c.Param("discuss_id")
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
	discussID, err := strconv.Atoi(discussIDString)
	if err != nil {
		log.Printf("search discuss error:%v", err)
		util.NormErr(c, 70013, "discuss_id非法")
		return
	}
	postID, userID, err := service.SearchPostAndUserByDiscussID(discussID) //得到帖子ID，被回复用户UID
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	var uDiscuss model.DiscussInfo
	uDiscuss.PostID = postID
	uDiscuss.Comment = comment
	uDiscuss.UserID = uUser.Id
	uDiscuss.ReplayID = discussID
	uDiscuss.ReplayUid = userID
	discussID, err = service.ReplayDiscuss(uDiscuss)
	if err != nil {
		log.Printf("search discuss error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.CreatDiscussRespSuccess(c, discussID)
}

func CheckReplay(c *gin.Context) {
	token := c.GetHeader("Authorization")
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
	u, err := service.CheckReplay(uUser.Id)
	if err != nil {
		log.Printf("search discuss error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.GetDiscussInfoSuccess(c, u)
}

func SubscribeComment(c *gin.Context) {
	postIDStr := c.PostForm("post_id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		log.Println(postIDStr)
		log.Println(err)
		util.NormErr(c, 70012, "post_id非法")
		return
	}
	rds.Subscribe(c, postID)
	util.RespOK(c)
}
