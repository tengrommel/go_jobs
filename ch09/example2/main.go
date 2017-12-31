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
	fmt.Println("Ok, connecting to redis!")
}
