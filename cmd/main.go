package main

import (
	"log"
	"meta_library/api"
	"meta_library/dao"
	"meta_library/dao/rds"
)

func main() {
	dao.InitDB()
	err := rds.InitRedis()
	if err != nil {
		log.Println(err)
		return
	}
	api.InitRouter()
}
