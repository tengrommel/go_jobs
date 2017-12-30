package main

import "fmt"

func test1()  {
	var a [10] int
	b := a
	b[0] = 200
	fmt.Println(a[0])
}

func main() {
	var a [10]int
	j := 9
	a[0] = 100
	a[j] = 200
	fmt.Println(a)

	for i:=0;i<len(a);i++{
		fmt.Println(a[i])
	}

	for index, item := range a{
		fmt.Printf("a[%d]=%d\n", index, item)
	}
	test1()
}
