package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age int
	Score float64
	next *Student
}

func trans(h *Student)  {
	for h!=nil{
		fmt.Println(*h)
		h = h.next //(*p).next go简化
	}

}

func main() {
	var head Student
	head.Name = "teng"
	head.Age = 34
	head.Score = 10

	var stu1 Student
	stu1.Name = "stu1"
	stu1.Age = 18
	stu1.Score = 90

	var stu2 Student
	stu2.Name = "stu2"
	stu2.Age = 18
	stu2.Score = 90

	var stu3 Student
	stu3.Name = "stu3"
	stu3.Age = 18
	stu3.Score = 90

	head.next = &stu1
	stu1.next = &stu2
	stu2.next = &stu3
	trans(&head)
}
