package main

import (
	"meta_library/api"
	"meta_library/dao"
)

func main() {
	dao.InitDB()
	api.InitRouter()
}
