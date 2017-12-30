package main

import "fmt"

type Student struct {
	Name string
	Age int
	Score float64
	left *Student
	right *Student
}

func trans(root *Student)  {
	if root==nil{
		return
	}
	fmt.Println(root)
	trans(root.right)
	trans(root.left)
}


func main() {
	var root *Student = new(Student)
	root.Name = "teng"
	root.Age = 34
	root.Score = 19.738947

	var left1 *Student = new(Student)
	left1.Name = "ni"
	left1.Age = 34
	left1.Score = 23.87439

	root.left = left1


	var right1 *Student = new(Student)
	right1.Name = "ni"
	right1.Age = 34
	right1.Score = 23.87439
	root.right = right1

	var right2 *Student = new(Student)
	right2.Name = "ni"
	right2.Age = 34
	right2.Score = 23.87439
	right1.right = right2

	trans(root)
}
