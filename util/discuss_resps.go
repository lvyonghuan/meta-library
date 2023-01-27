package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type creatDiscussRespSuccess struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
	Data   int    `json:"data"`
}

func CreatDiscussRespSuccess(c *gin.Context, data int) {
	response := creatCommentRespSuccess{
		Status: 10000,
		Info:   "success",
		Data:   data,
	}
	c.JSON(http.StatusOK, response)
}
