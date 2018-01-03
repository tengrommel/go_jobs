package main

import (
	"github.com/astaxie/beego"
	_ "github.com/tengrommel/go_jobs/ch14/SecKill/SecProxy/router"
	"github.com/astaxie/beego/logs"
	"fmt"
)

func initConfig() (err error) {
	redisAddr := beego.AppConfig.String("redis_addr")
	etcdAddr := beego.AppConfig.String("etcd_addr")

	logs.Debug("read config succ, redis addr:%v", redisAddr)
	logs.Debug("read config succ, etcd addr:%v", etcdAddr)

	secKillConf.etcdAddr = etcdAddr
	secKillConf.redisAddr = redisAddr
	if len(redisAddr)==0 ||len(etcdAddr) == 0{
		err = fmt.Errorf("init config failed, redis[%s] or etcd[%v] config is null", redisAddr, etcdAddr)
		return
	}
	return
}

func main()  {
	err := initConfig()
	if err != nil{
		panic(err)
		return
	}
	err = initSec()
	if err != nil{
		panic(err)
		return
	}
	beego.Run()
}
