package main

import (
	"github.com/tengrommel/go_jobs/ch11/logagent/tailf"
	"github.com/astaxie/beego/logs"
	"time"
	"github.com/tengrommel/go_jobs/ch11/logagent/kafka"
)

func serverRun() (err error) {
	for{
		msg := tailf.GetOneLine()
		err = sendToKafka(msg)
		if err != nil{
			logs.Error("send to kafka failed, err:%v", err)
			time.Sleep(time.Second)
			continue
		}
	}
	return
}

func sendToKafka(msg *tailf.TextMsg) (err error) {
	//logs.Debug("read msg:%s, topic:%s", msg.Msg, msg.Topic)
	kafka.SendToKafka(msg.Msg, msg.Topic)
	return
}