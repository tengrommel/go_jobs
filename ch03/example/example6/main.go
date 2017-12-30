package main

import (
	"fmt"
)

func Print(n int) {
	for index := 0; index < n+1; index++ {
		for j := 0; j < index; j++ {
			fmt.Printf("A")
		}
		fmt.Println()
	}
}

func main() {
	Print(6)
}
