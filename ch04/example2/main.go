package main

import "fmt"

func main() {
	var a []int
	a = append(a,10,20,383)
	a = append(a,a...)
	fmt.Println(a)
}
