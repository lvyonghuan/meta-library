package dao

import (
	"fmt"
	"meta_library/model"
)

func GetCommentList(bookID int) (u []model.CommentInfo, err error) {
	row, err := DB.Query("select * from comment where book_id=?", bookID)
	if err != nil {
		return
	}
	for row.Next() {
		var temp model.CommentInfo
		err = row.Scan(&temp.PostID, &temp.BookID, &temp.PublishTime, &temp.Content, &temp.UserID, &temp.Avatar, &temp.Nickname, &temp.PraiseCount, &temp.IsPraised, &temp.IsFocus)
		if err != nil {
			return
		}
		u = append(u, temp)
	}
	return
}

func CreatComment(u model.CommentInfo) (commentID int, err error) {
	res, err := DB.Exec("insert into comment(post_id,book_id,publish_time,content,user_id,avatar,nickname,praise_count,is_praised,is_focus) values (?,?,?,?,?,?,?,?,?,?)", u.PostID, u.BookID, u.PublishTime, u.Content, u.UserID, u.Avatar, u.Nickname, u.PraiseCount, u.IsPraised, u.IsFocus)
	if err != nil {
		return
	}
	commentID64, err := res.LastInsertId()
	commentID = int(commentID64)
	return
}

func RefreshComment(userID int, commentID int, content string) (err error) {
	var count int
	err = DB.QueryRow("SELECT COUNT(*) FROM comment WHERE post_id=? AND user_id=?", commentID, userID).Scan(&count)
	if err != nil {
		return err
	}
	if count != 1 {
		return fmt.Errorf("post_id and user_id not match")
	}
	_, err = DB.Exec("update comment set content=? where post_id=? and user_id=?", content, commentID, userID)
	return
}

func DeleteComment(userID int, commentID int) (err error) {
	var count int
	err = DB.QueryRow("SELECT COUNT(*) FROM comment WHERE post_id=? AND user_id=?", commentID, userID).Scan(&count)
	if err != nil {
		return err
	}
	if count != 1 {
		return fmt.Errorf("post_id and user_id not match")
	}
	_, err = DB.Exec("delete from comment where post_id=?", commentID)
	return
}
