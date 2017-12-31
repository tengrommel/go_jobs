package main

import (
	"fmt"
	"time"
)

func calc(taskChan chan int, resChan chan int)  {
	for v := range taskChan{
		flag := true
		for i:=2;i<v;i++ {
			if v%i==0{
				flag = false
				break
			}
		}
		if flag{
			resChan <- v
		}
	}
}


func main() {
	intChan := make(chan int, 1000)
	resultChan := make(chan int, 1000)
	for i:=0;i<1000;i++ {
		intChan <-i
	}

	close(intChan)

	for i:=0;i<8;i++ {
		go calc(intChan, resultChan)
	}

	go func() {
		for v := range resultChan{
			fmt.Println(v)
		}
	}()

	time.Sleep(time.Second*10)
}
