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

func GetCommentList(c *gin.Context) {
	bookIDString := c.Param("book_id")
	bookID, err := strconv.Atoi(bookIDString)
	if err != nil {
		log.Printf("search book error:%v", err)
		util.NormErr(c, 70002, "book_id非法")
		return
	}
	u, err := service.GetCommentList(bookID)
	if err != nil {
		log.Printf("search comment error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.CommentRespSuccess(c, u)
}

func CreatComment(c *gin.Context) {
	token := c.GetHeader("Authorization")
	bookIDString := c.Param("book_id")
	content := c.PostForm("content")
	bookID, err := strconv.Atoi(bookIDString)
	if err != nil {
		log.Printf("search book error:%v", err)
		util.NormErr(c, 70002, "book_id非法")
		return
	}
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
	uUser, err := service.SearchUserByUserName(username)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	var uComment model.CommentInfo
	uComment.BookID = bookID
	uComment.PublishTime = tool.FormatTime()
	uComment.Content = content
	uComment.UserID = uUser.Id
	uComment.Avatar = uUser.Avatar
	uComment.Nickname = uUser.Nickname
	uComment.PostID, err = service.CreatComment(uComment)
	if err != nil {
		log.Printf("search comment error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.CreatCommentRespSuccess(c, uComment.PostID)
}

func DeleteComment(c *gin.Context) {
	token := c.GetHeader("Authorization")
	commentIDString := c.Param("comment_id")
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
	commentID, err := strconv.Atoi(commentIDString)
	if err != nil {
		log.Printf("search comment error:%v", err)
		util.NormErr(c, 80002, "comment_id非法")
		return
	}
	err = service.DeleteComment(u.Id, commentID, u.IsAdministrator)
	if err != nil {
		if err.Error() == "post_id and user_id not match" {
			util.NormErr(c, 80004, "用户无权限删除此书评")
			return
		}
		log.Printf("search comment error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.RespOK(c)
}

func RefreshComment(c *gin.Context) {
	token := c.GetHeader("Authorization")
	commentIDString := c.Param("comment_id")
	content := c.PostForm("content")
	commentID, err := strconv.Atoi(commentIDString)
	if err != nil {
		log.Printf("search comment error:%v", err)
		util.NormErr(c, 80002, "comment_id非法")
		return
	}
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
	//验证更新操作者与发布者的信息的一致性
	u, err := service.SearchUserByUserName(username)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	err = service.RefreshComment(u.Id, commentID, content)
	if err != nil {
		if err.Error() == "post_id and user_id not match" {
			util.NormErr(c, 80003, "用户无权限更改此书评")
			return
		}
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.RespOK(c)
}
