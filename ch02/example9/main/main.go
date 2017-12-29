package main

import (
	"fmt"
)

func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func swap1(a int, b int) (int, int) {
	return b, a
}

func test() {
	var a = 100
	fmt.Println(a)

	for i := 0; i < 100; i++ {
		var b = i * 2
		fmt.Println(b)
	}
	// fmt.Println(b)
}

func main() {
	first := 100
	second := 200
	swap(&first, &second)
	fmt.Println("first=", first)
	fmt.Println("second=", second)

	swap1(first, second)
	fmt.Println("first=", first)
	fmt.Println("second=", second)
	test()
}
