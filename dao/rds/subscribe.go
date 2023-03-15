package rds

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
)

func Subscribe(c *gin.Context, postID int) {
	pubSub := Rds.Subscribe(c, strconv.Itoa(postID))
	defer func(pubSub *redis.PubSub) {
		err := pubSub.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}(pubSub)
}

func Publish(c *gin.Context, postID int, message string) {
	err := Rds.Publish(c, strconv.Itoa(postID), message)
	if err != nil {
		log.Println(err)
		return
	}
}
