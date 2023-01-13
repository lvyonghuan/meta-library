package service

import (
	"meta_library/dao"
	"meta_library/model"
)

func SearchUserByUserName(username string) (u model.UserInfo, err error) {
	u, err = dao.SearchUserByUserName(username)
	return u, err
}

func InsertUser(u model.UserInfo) (err error) {
	err = dao.InsertUser(u)
	return err
}

func ChangePasswordByUsername(username string, newPassword string) (err error) {
	err = dao.ChangePasswordByUsername(username, newPassword)
	return err
}

func SearchUserByUserId(id int) (u model.UserInfo, err error) {
	u, err = dao.SearchUserByUserId(id)
	return u, err
}

func ChangeUserInfo(u model.UserInfo) (err error) {
	err = dao.ChangeUserInfo(u)
	return err
}
