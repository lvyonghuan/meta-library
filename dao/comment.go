package dao

import (
	"fmt"
	"meta_library/model"
)

func GetCommentList(bookID int) (u []model.CommentInfo, err error) {
	err = DB.Table("comment").Where("book_id=?", bookID).Find(u).Error
	return
	//row, err := DB.Query("select * from comment where book_id=?", bookID)
	//if err != nil {
	//	return
	//}
	//for row.Next() {
	//	var temp model.CommentInfo
	//	err = row.Scan(&temp.PostID, &temp.BookID, &temp.PublishTime, &temp.Content, &temp.UserID, &temp.Avatar, &temp.Nickname, &temp.PraiseCount, &temp.IsPraised, &temp.IsFocus)
	//	if err != nil {
	//		return
	//	}
	//	u = append(u, temp)
	//}
	//return
}

func CreatComment(u model.CommentInfo) (commentID int, err error) {
	if err := DB.Table("comment").Create(&u).Error; err != nil {
		return -1, err
	}
	return u.PostID, nil
	//res, err := DB.Exec("insert into comment(post_id,book_id,publish_time,content,user_id,avatar,nickname,praise_count,is_praised,is_focus) values (?,?,?,?,?,?,?,?,?,?)", u.PostID, u.BookID, u.PublishTime, u.Content, u.UserID, u.Avatar, u.Nickname, u.PraiseCount, u.IsPraised, u.IsFocus)
	//if err != nil {
	//	return
	//}
	//commentID64, err := res.LastInsertId()
	//commentID = int(commentID64)
	//return
}

func RefreshComment(userID int, commentID int, content string) (err error) {
	var temp model.CommentInfo
	temp.PostID = commentID
	temp.UserID = userID
	temp.Content = content
	err = DB.Table("comment").Model(temp).Updates(temp).Error
	return
	//var count int
	//err = DB.QueryRow("SELECT COUNT(*) FROM comment WHERE post_id=? AND user_id=?", commentID, userID).Scan(&count)
	//if err != nil {
	//	return err
	//}
	//if count != 1 {
	//	return fmt.Errorf("post_id and user_id not match")
	//}
	//_, err = DB.Exec("update comment set content=? where post_id=? and user_id=?", content, commentID, userID)
	//return
}

func DeleteComment(userID int, commentID int, isAdministrator bool) (err error) {
	var comment model.CommentInfo
	//权限检查
	if result := DB.Table("comment").Where("id = ? AND user_id = ?", commentID, userID).First(&comment); result.Error != nil {
		return result.Error
	} else if !isAdministrator && comment.UserID != userID {
		return fmt.Errorf("you are not allowed to delete this comment")
	}
	//删除操作
	if result := DB.Table("comment").Delete(&comment, commentID); result.Error != nil {
		return result.Error
	}
	return nil

	//var count int
	//err = DB.QueryRow("SELECT COUNT(*) FROM comment WHERE post_id=? AND user_id=?", commentID, userID).Scan(&count)
	//if err != nil {
	//	return err
	//}
	//if count != 1 && !isAdministrator {
	//	return fmt.Errorf("post_id and user_id not match")
	//}
	//_, err = DB.Exec("delete from comment where post_id=?", commentID)
	//return
}
