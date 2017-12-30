package main

import (
	"fmt"
	"sort"
)

func testMapSort()  {
	var a map[int]int
	a = make(map[int]int)
	a[9]=0
	a[5]=0
	a[4]=0
	a[1]=3

	var keys []int
	for k := range a{
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, v :=range keys{
		fmt.Println(v, a[v])
	}
}

func revalsMap()  {
	var a map[int]string
	var b map[string]int
	a = make(map[int]string)
	b = make(map[string]int)

	a[123] = "123"
	a[234] = "234"

	for k, v := range a {
		b[v]=k
	}

	fmt.Println(b)
}

func main() {
	testMapSort()
	revalsMap()
}
