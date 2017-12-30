package main

import "fmt"

func testSlice()  {
	var slice []int
	var arr [5]int = [...]int{1,2,3,4,5}

	slice = arr[2:5]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = slice[0:1]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}

func main() {
	testSlice()
}
