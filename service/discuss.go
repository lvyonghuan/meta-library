package service

import (
	"meta_library/dao"
	"meta_library/model"
)

func CreateDiscuss(u model.DiscussInfo) (discussID int, err error) {
	discussID, err = dao.CreateDiscuss(u)
	return
}
