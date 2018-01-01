package main

import "fmt"

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
}
