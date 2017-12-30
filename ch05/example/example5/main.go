package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"name"`
	Age int	`json:"age"`
	Score int `json:"score"`
}

func main() {
	var stu Student = Student{
		Name:"teng",
		Age: 43,
		Score:56,
	}
	data, err := json.Marshal(stu)
	if err != nil{
		fmt.Println("json encode stu failed:", err)
	}
	fmt.Println(string(data))
}
