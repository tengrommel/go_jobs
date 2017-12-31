package main

import "fmt"

type student struct {
	name string
}

func main() {
	var stuChan chan *student//管道需要初始化
	stuChan = make(chan *student, 10)
	stu := student{name:"stud01"}
	stuChan <- &stu

	var stu01 *student
	stu01 =<- stuChan
	fmt.Println(stu01)
}
