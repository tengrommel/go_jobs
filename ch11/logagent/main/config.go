package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"errors"
)

var (
	appConfig *Config
)


type Config struct {
	logLevel string
	logPath string

	collectConf []CollectConf
}

type CollectConf struct {
	logPath string
	topic string
}

func loadCollectConf(conf config.Configer) (err error) {
	var cc CollectConf
	cc.logPath = conf.String("collect::log_path")
	if len(cc.logPath) == 0 {
		err = errors.New("")
		return
	}
	cc.topic = conf.String("collect::topic")
	if len(cc.logPath) == 0 {
		err = errors.New("")
		return
	}
	appConfig.collectConf = append(appConfig.collectConf, cc)
	return
}


func loadConf(confType, filename string) (err error) {
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}

	appConfig = &Config{}
	appConfig.logLevel = conf.String("logs::log_level")
	if len(appConfig.logLevel) == 0{
		appConfig.logLevel = "debug"
	}
	appConfig.logPath = conf.String("logs::log_path")
	if len(appConfig.logPath) == 0{
		appConfig.logPath = "./logs"
	}
	err = loadCollectConf(conf)
	if err != nil{
		fmt.Println("load collect conf failed, err:%v\n", err)
		return
	}
	return
}
