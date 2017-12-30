package main

import "fmt"

func test()  {
	defer func() {
		if err := recover(); err != nil{
			fmt.Println(err)
		}
	}()
	b := 0
	a := 100/b
	fmt.Println(a)
	return
}

func main() {
	test()
	var a []int
	a = append(a,10,20,383)
	a = append(a,a...)
	fmt.Println(a)
}
