package main

import "fmt"

type op_func func(int, int) int

func add(a, b int) int {
	return a + b
}

func operator(op op_func, a, b int) int {
	return op(a, b)
}

func main() {
	c := add
	sum := operator(c, 100, 299)
	fmt.Println(sum)
}
