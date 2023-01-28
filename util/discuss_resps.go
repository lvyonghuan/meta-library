package util

import (
	"github.com/gin-gonic/gin"
	"meta_library/model"
	"net/http"
)

type discussRespSuccess struct {
	Status   int                 `json:"status"`
	Info     string              `json:"info"`
	Comments []model.DiscussInfo `json:"comments"`
}

func GetDiscussInfoSuccess(c *gin.Context, u []model.DiscussInfo) {
	response := discussRespSuccess{
		Status:   10000,
		Info:     "success",
		Comments: u,
	}
	c.JSON(http.StatusOK, response)
}

func CreatDiscussRespSuccess(c *gin.Context, data int) {
	response := creatCommentRespSuccess{
		Status: 10000,
		Info:   "success",
		Data:   data,
	}
	c.JSON(http.StatusOK, response)
}
