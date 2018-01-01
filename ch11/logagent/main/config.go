package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"errors"
	"github.com/tengrommel/go_jobs/ch11/logagent/tailf"
)

var (
	appConfig *Config
)


type Config struct {
	logLevel string
	logPath string

	chanSize int
	collectConf []tailf.CollectConf
}



func loadCollectConf(conf config.Configer) (err error) {
	var cc tailf.CollectConf
	cc.LogPath = conf.String("collect::log_path")
	if len(cc.LogPath) == 0 {
		err = errors.New("")
		return
	}
	cc.Topic = conf.String("collect::topic")
	if len(cc.LogPath) == 0 {
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

	appConfig.chanSize, err = conf.Int("collect::chan_size")
	if err != nil{
		appConfig.chanSize=100
	}
	err = loadCollectConf(conf)
	if err != nil{
		fmt.Println("load collect conf failed, err:%v\n", err)
		return
	}
	return
}
