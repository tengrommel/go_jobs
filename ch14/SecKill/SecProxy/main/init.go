package main

import "github.com/astaxie/beego/logs"

func initRedis() (err error) {
	return
}

func initEtcd() (err error) {
	return
}

func initSec() (err error) {

	err = initRedis()
	if err != nil{
		logs.Error("init redis failed, err:%v", err)
		return
	}

	err = initEtcd()
	if err != nil{
		logs.Error("init etcd failed, err:%v", err)
		return
	}

	logs.Info("init sec succ")
	return
}
