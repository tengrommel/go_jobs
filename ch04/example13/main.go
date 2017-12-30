package main

import (
	"math/rand"
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex

func testMap()  {
	var a map[int]int
	a = make(map[int]int, 5)
	a[2]=1
	a[8]=8
	a[1]=5
	a[4]=4
	a[5]=1
	for i := 0; i<2;i++{
		go func(b map[int]int) {
			lock.Lock()
			b[9] = rand.Intn(100)
			lock.Unlock()
		}(a)
	}
	lock.Lock()
	fmt.Println(a)
	lock.Unlock()

	time.Sleep(time.Second)
}

func main() {
	testMap()


}
