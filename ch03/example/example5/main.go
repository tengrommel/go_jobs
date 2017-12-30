package main

import (
	"fmt"
)

func main() {
	var a int = 10
	switch a {
	case 0:
		fmt.Println("a is equal 0")
	case 10:
		fmt.Println("a is equal 10")
		fallthrough
	default:
		fmt.Println("a is equal default")
	}
}
