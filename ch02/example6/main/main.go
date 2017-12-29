package main

import (
	"fmt"
	"time"
)

const (
	Male   = 1
	Female = 2
)

func main() {
	for {
		second := time.Now().Unix()
		if second%Female == 0 {
			fmt.Println("female")
		} else {
			fmt.Println("male")
		}
		time.Sleep(1000 * time.Millisecond)
	}

}
