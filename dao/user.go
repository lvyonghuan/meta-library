package dao

import "meta_library/model"

func SearchUserByUserName(name string) (u model.UserInfo, err error) { //查找用户名
	row := DB.QueryRow("select * from user where username = ?", name)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.Id, &u.UserName, &u.PassWord, &u.Nickname, &u.Gender, &u.QQ, &u.Birthday, &u.Email, &u.Avatar, &u.Introduction, &u.Phone)
	return
}

func InsertUser(u model.UserInfo) (err error) { //注册，将用户信息填入数据库
	_, err = DB.Exec("insert into user(id,username,password,nickname,gender,qq,birthday,email,avatar,introduction,phone) values (?,?,?,?,?,?,?,?,?,?,?,?,?)", u.Id, u.UserName, u.PassWord, u.Nickname, u.Gender, u.QQ, u.Birthday, u.Email, u.Avatar, u.Introduction, u.Phone)
	return err
}

func ChangePasswordByUsername(usernmae string, newPassword string) (err error) { //改密码
	_, err = DB.Exec("update user set password=? where username=?", newPassword, usernmae)
	return err
}

func SearchUserByUserId(id int) (u model.UserInfo, err error) { //根据id查用户
	row := DB.QueryRow("select * from user where id = ?", id)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.Id, &u.UserName, &u.PassWord, &u.Nickname, &u.Gender, &u.QQ, &u.Birthday, &u.Email, &u.Avatar, &u.Introduction, &u.Phone)
	return
}

func ChangeUserInfo(u model.UserInfo) (err error) { //查询用户信息
	_, err = DB.Exec("update user set nickname=?,avatar=?,introduction=?,phone=?,qq=?,gender=?,email=?,birthday=? where username=?", u.Nickname, u.Avatar, u.Introduction, u.Phone, u.QQ, u.Gender, u.Email, u.Birthday, u.UserName)
	return err
}
