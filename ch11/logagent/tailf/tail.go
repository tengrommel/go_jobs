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

type TextMsg struct {
	Msg string
	Topic string
}

type TailObjMgr struct {
	tailObjs []*TailObj
	msgChan chan *TextMsg
}

var (
	tailObjMgr* TailObjMgr
)

func InitTail(conf []CollectConf, chanSize int) (err error) {
	if len(conf)==0{
		err = fmt.Errorf("invalid config")
		return
	}
	tailObjMgr = &TailObjMgr{
		msgChan: make(chan*TextMsg, chanSize),
	}
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
