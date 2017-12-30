package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("HELLO")
	time.Sleep(time.Second)
	main()
}
