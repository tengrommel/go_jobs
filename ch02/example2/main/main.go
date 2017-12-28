package main

import (
	"fmt"

	"github.com/tengrommel/go_jobs/ch02/example2/add"
)

func main() {
	add.Test()
	fmt.Println("Name=", add.Name)
	fmt.Println("Age=", add.Age)
}
