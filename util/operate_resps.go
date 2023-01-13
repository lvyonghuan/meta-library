package util

import (
	"github.com/gin-gonic/gin"
	"meta_library/model"
	"net/http"
)

type respCollectSuccess struct {
	Status int             `json:"status"`
	Info   string          `json:"info"`
	Data   collectInfoList `json:"data"`
}

type collectInfoList struct {
	Collections []collectBookInfoList `json:"collections"`
}

type collectBookInfoList struct {
	BookId      int    `json:"book_id"`
	Name        string `json:"name"`
	PublishTime string `json:"publish_time"`
	Link        string `json:"link"`
}

func RespCollectSuccess(c *gin.Context, u []model.BookInfo) {
	var temp []collectBookInfoList
	for _, book := range u {
		temp = append(temp, collectBookInfoList{
			BookId:      book.BookId,
			Name:        book.Name,
			PublishTime: book.PublishTime,
			Link:        book.Link,
		})
	}
	response := respCollectSuccess{
		Status: 10000,
		Info:   "success",
		Data:   collectInfoList{temp},
	}
	c.JSON(http.StatusOK, response)
}
