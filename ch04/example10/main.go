package main

import (
	"sort"
	"fmt"
)

func testIntSort()  {
	var a = [...]int{1,4,3,2,9}
	sort.Ints(a[:])
	fmt.Println(a)
}

func testStringSort()  {
	var a = [...]string{"abc", "efg", "b", "A", "eeee"}
	sort.Strings(a[:])
	fmt.Println(a)
}

func testFloatSort()  {
	var a = [...]float64{2.3,5.4,1.1}
	sort.Float64s(a[:])
	fmt.Println(a)
}

// 必须为排序后
func testIntSearch()  {
	var a = [...]int{1,6,4,0,2}
	index := sort.SearchInts(a[:], 6)
	fmt.Println(index)
}

func main() {
	testIntSort()
	testStringSort()
	testFloatSort()
	testIntSearch()
}
