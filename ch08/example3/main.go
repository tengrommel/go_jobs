package main

import "fmt"

func send (ch chan<- int, lock chan bool)  {
	for i:=0;i<10;i++ {
		ch <- i
	}
	close(ch)
	lock <- true
}

func recv (ch <-chan int, lock chan bool)  {
	for  {
		v, ok := <- ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
	lock<- true
}
func main() {
	ch := make(chan int, 10)
	lock := make(chan bool, 2)
	go send(ch, lock)
	go recv(ch, lock)

	var total = 0
	for _ = range lock{
		total++
		if total == 2{
			break
		}
	}
}
