package main

import "fmt"

type Student struct {
	Name string
	Age int
	score float64
}

func main() {
	var stu Student
	stu.Age = 18
	stu.Name = "h"
	stu.score = 80

	var stu1 *Student = &Student{
		Age: 35,
		Name: "teng",

	}

	fmt.Println(stu)
	fmt.Println(stu1)
	fmt.Printf("%p, %p, %p", &stu.Name, &stu.Age, &stu.score)
}
