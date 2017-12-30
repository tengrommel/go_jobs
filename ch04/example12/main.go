package main

import "fmt"

func testMapSort()  {
	var a map[int]int
	a = make(map[int]int)
	a[9]=0
	a[5]=0
	a[4]=0
	a[1]=3

	for k, v := range a{
		fmt.Println(k, v)
	}
}

func main() {
	testMapSort()
}
