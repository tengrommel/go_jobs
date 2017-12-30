package main

import "fmt"

type student struct {
	Name string
	Age int
	Score float64
}

func main() {
	var str = "teng 45 90.3"
	var stu student
	fmt.Sscanf(str, "%s %d %f", &stu.Name, &stu.Age, &stu.Score)
	fmt.Println(stu)
}
