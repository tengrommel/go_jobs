package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int
	ch = make(chan int, 10)
	ch2 := make(chan int, 10)

	go func() {
		var i int
		for  {
			ch <- i
			time.Sleep(time.Second)
			ch2 <- i*2
			time.Sleep(time.Second)
			i++
		}
	}()

	for  {
		select {
		case v := <- ch:
			fmt.Println(v)
		case v:= <-ch2:
			fmt.Println(v)
		default:
			fmt.Println("get data timeout!")
			time.Sleep(time.Second)
		}
	}

}
