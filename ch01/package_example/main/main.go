package main

import (
	"fmt"

	a "github.com/tengrommel/go_jobs/ch01/package_example/calc"
)

func main() {
	sum := a.Add(100, 300)
	sub := a.Sub(100, 300)
	fmt.Println("sum=", sum)
	fmt.Println("sub=", sub)
}
