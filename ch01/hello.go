package main

import (
	"time"
)

func add(a int, b int) int {
	var sum int
	sum = a + b
	return sum
}

func main() {
	// var c int
	// c = add(100, 200)
	//go testGoroutine(300, 300)
	// 有可能不输出
	for i := 0; i < 100; i++ {
		go testPrint(i)
	}
	// fmt.Println("add(100, 200) =    ", c)
	time.Sleep(time.Second)
}
