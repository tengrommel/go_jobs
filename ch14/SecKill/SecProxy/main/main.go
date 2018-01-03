package main

import (
	"github.com/astaxie/beego"
	_ "github.com/tengrommel/go_jobs/ch14/SecKill/SecProxy/router"
	"github.com/astaxie/beego/logs"
	"fmt"
)


func main()  {
	err := initConfig()
	if err != nil {
		panic(err)
		return
	}

	err = initSec()
	if err != nil {
		panic(err)
		return
	}

	beego.Run()
}
