package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	file, err := os.Open("")
	if err != nil{
		fmt.Println("read file error", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	str, err := reader.ReadString('\n')
	if err!= nil{
		fmt.Println()
	}
	fmt.Println(str)
}
