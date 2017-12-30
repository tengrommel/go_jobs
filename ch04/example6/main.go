package main

import "fmt"

func main() {
	var a [10]int
	j := 10
	a[0] = 100
	a[j] = 200
	fmt.Println(a)
}
