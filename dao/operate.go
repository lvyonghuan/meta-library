package dao

import (
	"meta_library/model"
)

func PraiseComment(commentID int) (err error) {
	_, err = DB.Exec("update comment set praise_count=praise_count+1 where post_id=?", commentID)
	return
}

func GetCollectList(userID int) (u []model.BookInfo, err error) {
	row, err := DB.Query("select book.* from book join star on star.bookId=book.book_id where star.Id=?", userID)
	if err != nil {
		return
	}
	for row.Next() {
		var temp model.BookInfo
		err = row.Scan(&temp.BookId, &temp.Name, &temp.IsStar, &temp.Author, &temp.CommentNum, &temp.Score, &temp.Cover, &temp.PublishTime, &temp.Link, &temp.Label)
		if err != nil {
			return
		}
		u = append(u, temp)
	}
	return
}

func Focus(followerID int, followeeID int) (err error) {
	_, err = DB.Exec("insert into follow(follower_id,followee_id) values (?,?)", followerID, followeeID)
	return
}
