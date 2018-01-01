package tailf

import (
	"github.com/hpcloud/tail"
	"fmt"
)

type CollectConf struct {
	LogPath string
	Topic string
}

type TailObj struct {
	tail *tail.Tail
	conf CollectConf
}

type TailObjMgr struct {
	tailObjs []*TailObj
}

var (
	tailObjMgr* TailObjMgr
)

func InitTail(conf []CollectConf) (err error) {
	if len(conf)==0{
		err = fmt.Errorf("invalid config")
		return
	}
	tailObjMgr = &TailObjMgr{}
	for _, v := range conf{
		obj := &TailObj{
			conf:v,
		}
		tails, errTail := tail.TailFile(v.LogPath, tail.Config{
			ReOpen:    true,
			Follow:    true,
			Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
			MustExist: false,
			Poll:      true,
		})
		if err != nil {
			err = errTail
			return
		}
		obj.tail = tails
		tailObjMgr.tailObjs=append(tailObjMgr.tailObjs, obj)
	}
	return
}
