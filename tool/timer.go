package tool

import (
	"log"
	"meta_library/dao"
	"time"
)

func DeleteSessionTimer(sessionID string) (err error) {
	timer := time.NewTimer(time.Duration(3600) * time.Second)
	<-timer.C //3600秒之前，管道阻塞。之后执行删除程序。
	err = dao.DeleteSession(sessionID)
	if err != nil {
		log.Printf("search timer error:%v", err)
		return
	}
	return
}
