package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"meta_library/service"
	"meta_library/tool"
	"meta_library/util"
	"strconv"
)

func GetBookList(c *gin.Context) {
	u, err := service.GetBookList()
	if err != nil {
		log.Printf("search book error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.BookListRespSuccess(c, u)
}

func SearchBookInfo(c *gin.Context) {
	token := c.GetHeader("Authorization")
	bookName := c.Query("book_name")
	uBook, err := service.SearchBook(bookName)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			util.NormErr(c, 70001, "数据库中没有该书籍")
			return
		}
		log.Printf("search book error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	for token != "" {
		isExist, username, err := tool.TokenExpired([]byte("114"), token)
		if err != nil {
			log.Printf("search user error:%v", err)
			util.NormErr(c, 60100, "token错误")
			break
		}
		if !isExist {
			util.NormErr(c, 60102, "token已过期")
			break
		}
		uUser, err := service.SearchUserByUserName(username)
		if err != nil {
			log.Printf("search user error:%v", err)
			util.RsepInternalErr(c)
			break
		}
		isStar, err := service.SearchUserStar(uUser.Id, uBook.BookId)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				break
			}
			log.Printf("search user error:%v", err)
			util.RsepInternalErr(c)
			break
		}
		if isStar {
			uBook.IsStar = true
			break
		}
		break
	}
	util.BookRespSuccess(c, uBook)
}

func StarBook(c *gin.Context) { //注：疑似还得加个取消收藏，写完了再说。另外待加入的功能：判断书籍id，判断是否已收藏。
	token := c.GetHeader("Authorization")
	bookIdString := c.Query("book_id") //文档里好像没写位置，我擅自用query了
	isExist, username, err := tool.TokenExpired([]byte("114"), token)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 60100, "token错误")
		return
	}
	if !isExist {
		util.NormErr(c, 60102, "token过期")
		return
	}
	uUser, err := service.SearchUserByUserName(username)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	bookId, err := strconv.Atoi(bookIdString)
	if err != nil {
		log.Printf("search book error:%v", err)
		util.NormErr(c, 70002, "book_id非法")
		return
	}
	err = service.StarBook(uUser.Id, bookId)
	if err != nil {
		log.Printf("search book error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.RespOK(c)
}

func GetSameLabelBook(c *gin.Context) {
	label := c.Query("label")
	u, err := service.GetSameLabelBook(label)
	if err != nil {
		log.Printf("search book error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.BookListRespSuccess(c, u)
}
