package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err!= nil{
		fmt.Println("conn redis failed,", err)
		return
	}
	defer c.Close()
	_, err = c.Do("Set", "abc", 100)
	if err != nil{
		fmt.Println(err)
		return
	}
	r,err := redis.Int(c.Do("Get", "abc"))
	if err!= nil{
		fmt.Println("gett abc failed,", err)
	}
	fmt.Println(r)
}
