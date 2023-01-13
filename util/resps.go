package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type respTemplate struct {
	Status int    `json:"status"` //状态码
	Info   string `json:"info"`   //访问信息
}

//结构体↑

var ok = respTemplate{ //成功操作
	Status: 10000,
	Info:   "success",
}

var ParamError = respTemplate{ //错误操作
	Status: 30000,
	Info:   "params error",
}

var InternalError = respTemplate{ //访问错误
	Status: 50000,
	Info:   "internal error",
}

//模型↑ 相应模块↓

func RespOK(c *gin.Context) { //正确
	c.JSON(http.StatusOK, ok)
}
func RespParamErr(c *gin.Context) { //错误
	c.JSON(http.StatusBadRequest, ParamError)
}

func RsepInternalErr(c *gin.Context) { //连接错误
	c.JSON(http.StatusInternalServerError, InternalError)
}

func NormErr(c *gin.Context, status int, info string) { //其他错误
	c.JSON(http.StatusBadRequest, gin.H{
		"status": status,
		"info":   info,
	})
}
