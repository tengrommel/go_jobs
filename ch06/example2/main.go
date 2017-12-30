package main

import "fmt"

func Test(a interface{})  {
	b := a.(int)
	b += 3
	fmt.Println(b)
}

func main() {
	var b int
	Test(b)
}