package main

import (
	"math/rand"
	"fmt"
	"sync"
	"time"
	"sync/atomic"
)

var lock sync.Mutex
var rwlock sync.RWMutex

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

func testRWLock()  {
	var a map[int]int
	a = make(map[int]int, 5)
	var count int64
	a[2]=1
	a[8]=8
	a[1]=5
	a[4]=4
	a[5]=1
	for i := 0; i<2;i++{
		go func(b map[int]int) {
			rwlock.Lock()
			b[9] = rand.Intn(100)
			time.Sleep(10*time.Millisecond)
			rwlock.Unlock()
		}(a)
	}
	for i := 0; i<100;i++{
		go func(b map[int]int) {
			for {
				rwlock.RLock()
				time.Sleep(time.Millisecond)
				rwlock.RUnlock()
				atomic.AddInt64(&count, 1)
			}
		}(a)
	}

	time.Sleep(10*time.Second)
	fmt.Println(atomic.LoadInt64(&count))
}

func main() {
	testMap()
	testRWLock()

}
