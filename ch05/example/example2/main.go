package main

import "fmt"

type Student struct {
	Name string
	Age int
	Score float64
	next *Student
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

	var p *Student = &head
	for p!=nil{
		fmt.Println(*p)
		p = p.next //(*p).next go简化
	}
}
