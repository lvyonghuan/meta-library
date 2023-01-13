package util

import (
	"github.com/gin-gonic/gin"
	"meta_library/model"
	"net/http"
)

// List的结构体
type bookListResp struct {
	Status int          `json:"status"`
	Info   string       `json:"info"`
	Data   bookInfoList `json:"data"`
}
type bookInfoList struct {
	Books []model.BookInfo `json:"books"`
}

// 单本模型
type bookResp struct {
	Status int      `json:"status"`
	Info   string   `json:"info"`
	Data   bookInfo `json:"data"`
}

type bookInfo struct {
	Book model.BookInfo `json:"book"`
}

func BookListRespSuccess(c *gin.Context, u []model.BookInfo) {
	response := bookListResp{
		Status: 10000,
		Info:   "success",
		Data:   bookInfoList{u},
	}
	c.JSON(http.StatusOK, response)
}

func BookRespSuccess(c *gin.Context, u model.BookInfo) {
	response := bookResp{
		Status: 10000,
		Info:   "success",
		Data:   bookInfo{u},
	}
	c.JSON(http.StatusOK, response)
}
