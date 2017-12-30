package main

import "fmt"

func isort(a []int)  {
	for i:=1;i<len(a);i++{
		for j:=i;j>0;j-- {
			if a[j] > a[j-1]{
				break
			}
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}

func main() {
	b := [10]int{8,7,5,4,2,19,10,9,3,11}
	isort(b[:])
	fmt.Println(b)
}
