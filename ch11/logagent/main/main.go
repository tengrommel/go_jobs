package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/tengrommel/go_jobs/ch11/logagent/tailf"
)

func main() {
	filename := "./conf/logagent.conf"
	err := loadConf("ini",filename)
	if err != nil{
		fmt.Printf("load conf failed, err:%v\n", err)
		panic("load conf failed")
		return
	}
	err = initLogger()
	if err!=nil{
		fmt.Println("load logger failer, err:%v\n", err)
		panic("load logger failed")
		return
	}


	logs.Debug("load conf success, config:%v", appConfig)

	err = tailf.InitTail(appConfig.collectConf, appConfig.chanSize)
	if err != nil{
		logs.Error("init tail failed, err:%v", err)
		return
	}
	logs.Debug("initialize success")
	err = serverRun()
	if err != nil{
		logs.Error("init tail failed, err:%v", err)
		return
	}
	logs.Info("program exited")
}
