package main

import "fmt"

type Student struct {
	Name string
	Age int
	Score int
	sex int
}

func (p Student) init(name string, age int, score int)  {
	p.Name = name
	p.Age = age
	p.Score = score
	fmt.Println(p)
}

func main() {
	var stu Student
	stu.init("stu", 34, 33)
	fmt.Println(stu)
}
