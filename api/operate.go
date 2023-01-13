package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"meta_library/service"
	"meta_library/tool"
	"meta_library/util"
	"strconv"
)

func Praise(c *gin.Context) {
	token := c.GetHeader("Authorization")
	//以下全为擅自决定（指字段位置）
	model := c.Query("model")
	targetIDString := c.Param("target_id")
	isExist, _, err := tool.TokenExpired([]byte("114"), token)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 600100, "token错误")
		return
	}
	if !isExist {
		util.NormErr(c, 600102, "token已过期")
		return
	}
	targetID, err := strconv.Atoi(targetIDString)
	if err != nil {
		log.Printf("search comment error:%v", err)
		util.NormErr(c, 80502, "target_id非法")
		return
	}
	if model == "1" { //1为评论相关操作
		err = service.PraiseComment(targetID)
		if err != nil {
			log.Printf("search operate error:%v", err)
			util.RsepInternalErr(c)
			return
		}
	} else {
		util.NormErr(c, 80501, "非法操作")
	}
	util.RespOK(c)
}

func GetCollectList(c *gin.Context) {
	token := c.GetHeader("Authorization")
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
	u, err := service.SearchUserByUserName(username)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	uList, err := service.GetCollectList(u.Id)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.RespCollectSuccess(c, uList)
}

func FocusUser(c *gin.Context) {
	token := c.GetHeader("Authorization")
	followeeUserIDString := c.Param("user_id")
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
	u, err := service.SearchUserByUserName(username)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	followeeId, err := strconv.Atoi(followeeUserIDString)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 60004, "UID非法")
		return
	}
	err = service.Focus(u.Id, followeeId)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.RespOK(c)
}
