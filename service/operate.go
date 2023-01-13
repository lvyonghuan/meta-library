package service

import (
	"meta_library/dao"
	"meta_library/model"
)

func PraiseComment(commentID int) (err error) {
	err = dao.PraiseComment(commentID)
	return
}

func GetCollectList(userID int) (u []model.BookInfo, err error) {
	u, err = dao.GetCollectList(userID)
	return
}

func Focus(followerID int, followeeID int) (err error) {
	err = dao.Focus(followerID, followeeID)
	return
}
