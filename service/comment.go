package service

import (
	"meta_library/dao"
	"meta_library/model"
)

func GetCommentList(bookID int) (u []model.CommentInfo, err error) {
	u, err = dao.GetCommentList(bookID)
	return
}

func CreatComment(u model.CommentInfo) (commentID int, err error) {
	commentID, err = dao.CreatComment(u)
	return
}

func RefreshComment(userID int, commentID int, content string) (err error) {
	err = dao.RefreshComment(userID, commentID, content)
	return err
}

func DeleteComment(userID int, commentID int) (err error) {
	err = dao.DeleteComment(userID, commentID)
	return
}
