package util

import (
	"github.com/gin-gonic/gin"
	"meta_library/model"
	"net/http"
)

type commentRespSuccess struct {
	Status   int                 `json:"status"`
	Info     string              `json:"info"`
	Comments []model.CommentInfo `json:"comments"`
}

type creatCommentRespSuccess struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
	Data   int    `json:"data"`
}

func CommentRespSuccess(c *gin.Context, u []model.CommentInfo) {
	response := commentRespSuccess{
		Status:   10000,
		Info:     "success",
		Comments: u,
	}
	c.JSON(http.StatusOK, response)
}

func CreatCommentRespSuccess(c *gin.Context, postID int) {
	response := creatCommentRespSuccess{
		Status: 10000,
		Info:   "success",
		Data:   postID,
	}
	c.JSON(http.StatusOK, response)
}
