package main

import (
	"fmt"
)

func testPipe() {
	pipe := make(chan int, 3)
	pipe <- 1
	pipe <- 2
	pipe <- 3
	pipe <- 4
	fmt.Println(len(pipe))
}
