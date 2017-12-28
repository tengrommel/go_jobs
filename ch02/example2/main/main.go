package main

import (
	"fmt"

	"github.com/tengrommel/go_jobs/ch02/example2/add"
	_ "github.com/tengrommel/go_jobs/ch02/example2/test"
)

func main() {
	fmt.Println("Name=", add.Name)
	fmt.Println("Age=", add.Age)
}
