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

func just(items ...interface{}) {
	for index, value := range items{
		switch value.(type) {
		case bool:
			fmt.Printf("%d params is bool, %v\n", index, value)
		case int:
			fmt.Printf("%d params is int, %v\n" ,index, value)
		case float32, float64:
			fmt.Printf("%d params is float, %v\n", index, value)
		case string:
			fmt.Printf("%d params is string, %v\n", index, value)
		case Student:
			fmt.Printf("%d params is student, %v\n", index, value)
		case *Student:
			fmt.Printf("%d params is student point, %v\n", index, value)
		}
	}
}

func main() {
	var b int
	Test(b)
	var a  Student
	Test(a)
	just(b, 2.3, "strubg", a, &a)
}