package main

import (
	"fmt"
)

type Student struct {
	Name string
	Sex string
}

func Test(a interface{})  {
	// 接口转化为具体的类.(int)
	b, ok := a.(int)
	if ok == false{
		fmt.Println("CONVERT FAILED")
		return
	}
	b += 3
	fmt.Println(b)
}

func main() {
	var b int
	Test(b)
	var a  Student
	Test(a)
}