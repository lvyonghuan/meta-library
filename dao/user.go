package dao

import (
	"errors"
	"gorm.io/gorm"
	"meta_library/model"
)

func SearchUserByUserName(name string) (u model.UserInfo, err error) { //查找用户名
	err = DB.Table("user").Where("username = ?", name).First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 没有找到记录
		return u, nil
	}
	return
}

func InsertUser(u model.UserInfo) (err error) { //注册，将用户信息填入数据库
	result := DB.Table("user").Create(&u)
	return result.Error
}

func ChangePasswordByUsername(userID int, newPassword string) (err error) { //改密码
	var user model.UserInfo
	user.Id = userID
	user.PassWord = newPassword
	DB.Model(&user).Table("user").Where("id", userID).Select("password").Updates(newPassword)
	return
	//_, err = DB.Exec("update user set password=? where username=?", newPassword, username)
	//return err
}

func SearchUserByUserId(id int) (u model.UserInfo, err error) { //根据id查用户
	err = DB.Table("user").Where("id = ?", id).First(&u).Error
	return
}

func ChangeUserInfo(u model.UserInfo) (err error) { //修改用户信息
	DB.Table("user").Model(&u).Updates(u)
	return
	//_, err = DB.Exec("update user set nickname=?,avatar=?,introduction=?,phone=?,qq=?,gender=?,email=?,birthday=? where username=?", u.Nickname, u.Avatar, u.Introduction, u.Phone, u.QQ, u.Gender, u.Email, u.Birthday, u.UserName)
	//return err
}

func LinkWithGithub(githubID int, uID int) (err error) {
	err = DB.Table("github_relate").Create(&model.GithubRelate{GithubID: githubID, UID: uID}).Error
	return
	//_, err = DB.Exec("insert into github_relate(github_id,uid) values (?,?)", githubID, uID)
	//return
}

func SearchGithubID(githubID int) (uid int, err error) {
	relate := model.GithubRelate{}
	err = DB.Table("github_relate").Where("github_id = ?", githubID).First(&relate).Error
	return relate.UID, err
	//row := DB.QueryRow("select * from github_relate where github_id = ?", githubID)
	//if err = row.Err(); row.Err() != nil {
	//	return
	//}
	//err = row.Scan(&githubID, &uid)
	//return uid, err
}

func StoreSession(sessionID string, value string) (err error) {
	err = DB.Table("session").Create(&model.Session{Session: sessionID, Token: value}).Error
	return
	//_, err = DB.Exec("insert into session(session,token) values (?,?)", sessionID, value)
	//return
}

func DeleteSession(sessionID string) (err error) {
	err = DB.Table("session").Where("session = ?", sessionID).Delete(&model.Session{}).Error
	return
	//_, err = DB.Exec("delete from session where session=?", sessionID)
	//return err
}

func SearchSessionByID(sessionID string) (token string, err error) {
	var session = model.Session{
		Session: sessionID,
		Token:   token,
	}
	err = DB.Table("session").Where("session = ?", sessionID).First(&session).Error
	return session.Token, err
	//row := DB.QueryRow("select * from session where session = ?", sessionID)
	//if err = row.Err(); row.Err() != nil {
	//	return
	//}
	//err = row.Scan(&sessionID, &token)
	//return token, err
}
