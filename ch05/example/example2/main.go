package main

import (
	"fmt"
	"math/rand"
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
	fmt.Println()
}

func insertTail(p *Student)  {
	var tail = p
	for i:=0;i<10;i++{
		var stu = Student{
			Name: fmt.Sprintf("stu%d", i),
			Age: rand.Intn(100),
			Score: rand.Float64() * 100,
		}
		tail.next = &stu
		tail = &stu
	}
}

func main() {
	var head Student
	head.Name = "teng"
	head.Age = 34
	head.Score = 10.09483948150125

	insertTail(&head)
	trans(&head)
}
