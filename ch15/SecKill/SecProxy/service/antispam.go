package service

import (
	"sync"
	"fmt"
)

var (
	secLimitMgr = &SecLimitMgr{
		UserLimitMap:make(map[int]*SecLimit, 10000),
	}
)

type SecLimitMgr struct {
	UserLimitMap map[int]*SecLimit
	lock sync.Mutex
}

func antiSpam(req *SecRequest) (err error) {
	secLimitMgr.lock.Lock()
	secLimit, ok := secLimitMgr.UserLimitMap[req.UserId]
	if !ok{
		secLimit = &SecLimit{}
		secLimitMgr.UserLimitMap[req.UserId]= secLimit
	}
	count := secLimit.Count(req.AccessTime.Unix())
	secLimitMgr.lock.Unlock()

	if count >= secKillConf.UserSecAccessLimit{
		err = fmt.Errorf("invalid request")
		return
	}
	return
}

type SecLimit struct {
	count int
	curTime int64
}

func (p *SecLimit)Count(nowTime int64) (curCount int) {
	if p.curTime != nowTime {
		p.count = 1
		p.curTime = nowTime
		curCount = p.count
		return
	}
	p.count++
	curCount = p.count
	return
}

func (p *SecLimit)Check(nowTime int64) int {
	if p.curTime != nowTime {
		return 0
	}
	return p.count
}