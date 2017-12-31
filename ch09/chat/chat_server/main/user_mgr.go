package main

import "github.com/tengrommel/go_jobs/ch09/chat/chat_server/model"

var (
	mgr *model.UserMgr
)

func initUserMgr() {
	mgr = model.NewUserMgr(pool)
}
