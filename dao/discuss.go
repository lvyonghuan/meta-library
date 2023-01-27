package dao

import "meta_library/model"

//两个功能：一个是发表评论，第二个回复评论
//这块完全没用经验，姑且先弄一下思路。发表评论直接挂靠在post_id下即可，同时获取user_id。
//回复需要post_id和replay_id两个数据，同时获取user_id。
//参考百度贴吧（先吐槽一下百度贴吧，多多少少有点屎山）。删除帖子的时候根据post_id进行全部删除，但是删评论不会删回复。

func CreateDiscuss(u model.DiscussInfo) (discussID int, err error) {
	res, err := DB.Exec("insert into discuss(discuss_id,post_id,replay_id,comment,user_id,star_num) values (?,?,?,?,?,?)", u.DiscussID, u.PostID, u.ReplayID, u.Comment, u.UserID, u.StarNum)
	if err != nil {
		return
	}
	discussID64, err := res.LastInsertId()
	discussID = int(discussID64)
	return
}
