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

func DeleteDiscuss(discussID int, userID int) (err error) {
	err = dao.DeleteDiscuss(discussID, userID)
	return
}