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

func DeleteDiscuss(discussID int, userID int, isAdministrator bool) (err error) {
	err = dao.DeleteDiscuss(discussID, userID, isAdministrator)
	return
}

func ReplayDiscuss(u model.DiscussInfo) (discussID int, err error) {
	discussID, err = dao.ReplayDiscuss(u)
	return
}

func SearchPostAndUserByDiscussID(discussID int) (postID int, userID int, err error) {
	postID, userID, err = dao.SearchPostAndUserByDiscussID(discussID)
	return
}

func CheckReplay(userID int) (u []model.DiscussInfo, err error) {
	u, err = dao.CheckReplay(userID)
	return
}
