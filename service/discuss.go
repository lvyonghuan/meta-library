package service

import (
	"meta_library/dao"
	"meta_library/model"
)

func CreateDiscuss(u model.DiscussInfo) (discussID int, err error) {
	discussID, err = dao.CreateDiscuss(u)
	return
}

func GetDiscussList(postID int) (u []model.DiscussInfo, err error) {
	u, err = dao.GetDiscussList(postID)
	return
}
