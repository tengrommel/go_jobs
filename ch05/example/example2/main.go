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

	head.next = &stu1
	trans(&head)
}
