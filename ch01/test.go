package main

import (
	"fmt"
)

func testGoroutine(a int, b int) {
	sum := a + b
	fmt.Println(sum)
}

func testPrint(a int) {
	fmt.Println(a)
}
