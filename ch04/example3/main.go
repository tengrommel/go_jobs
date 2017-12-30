package main

import "fmt"

func test()  {
	s1 := new([]int)
	fmt.Println(s1)
	s2 := make([]int, 10)
	fmt.Println(s2)
	// slice是空的
	(*s1)[0] = 100
	s2[0] = 100
	return
}

func main() {
	test()
}
