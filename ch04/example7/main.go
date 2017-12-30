package main

import "fmt"

func fab(n int)  {
	var a []int
	a = make([]int, n)
	a[0] = 1
	a[1] = 1

	for i:=2;i<n;i++{
		a[i] = a[i-1]+a[i-2]
	}
	for _, v := range a{
		fmt.Println(v)
	}

}

func main() {
	fab(50)
}
