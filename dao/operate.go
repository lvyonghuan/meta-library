package dao

import (
	"gorm.io/gorm"
	"meta_library/model"
)

func PraiseComment(commentID int) (err error) {
	err = DB.Model(&model.CommentInfo{}).Where("post_id = ?", commentID).Update("praise_count", gorm.Expr("praise_count + ?", 1)).Error
	return
	//_, err = DB.Exec("update comment set praise_count=praise_count+1 where post_id=?", commentID)
	//return
}

func PraiseDiscuss(discussID int) (err error) {
	err = DB.Model(&model.DiscussInfo{}).Where("discuss_id = ?", discussID).Update("praise_count", gorm.Expr("praise_count + ?", 1)).Error
	return
	//_, err = DB.Exec("update discuss set praise_count=praise_count+1 where discuss_id=?", discussID)
	//return
}

func GetCollectList(userID int) (u []model.BookInfo, err error) {
	err = DB.Model(&model.BookInfo{}).Joins("JOIN star on star.book_id = book.book_id").Where("star.Id = ?", userID).Find(&u).Error
	return
	//row, err := DB.Query("select book.* from book join star on star.bookId=book.book_id where star.Id=?", userID)
	//if err != nil {
	//	return
	//}
	//for row.Next() {
	//	var temp model.BookInfo
	//	err = row.Scan(&temp.BookId, &temp.Name, &temp.IsStar, &temp.Author, &temp.CommentNum, &temp.Score, &temp.Cover, &temp.PublishTime, &temp.Link, &temp.Label)
	//	if err != nil {
	//		return
	//	}
	//	u = append(u, temp)
	//}
	//return
}

func Focus(followerID int, followeeID int) (err error) {
	var temp model.Follow
	temp.FollowerID = followerID
	temp.FolloweeID = followeeID
	err = DB.Create(&temp).Error
	return
	//_, err = DB.Exec("insert into follow(follower_id,followee_id) values (?,?)", followerID, followeeID)
	//return
}
