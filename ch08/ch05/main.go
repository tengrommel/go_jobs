package main

import (
	"time"
	"fmt"
)

func main() {
	go func() {
		for  {
			t := time.NewTicker(time.Second)
			select {
			case <-t.C:
				fmt.Println("ok")
			}
			t.Stop()
		}
	}()

	time.Sleep(time.Second*3)
}
