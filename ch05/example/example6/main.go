package main

import (
	"time"
	"fmt"
)

type Cart struct {
	name string
	age int
}

type Train struct {
	Cart
	int
	start time.Time
}

func main() {
	var t Train
	t.name = "fuck"
	t.Cart.age = 34
	t.int = 200
	fmt.Println(t)
}
