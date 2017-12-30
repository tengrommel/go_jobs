package main

import "fmt"

func ssort(a []int)  {
	for i:=0;i<len(a);i++ {
		var min int = i
		for j:=i+1;j<len(a);j++ {
			if a[min]>a[j] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
}


func main() {
	b := [10]int{8,7,5,4,2,19,10,9,3,11}
	ssort(b[:])
	fmt.Println(b)
}
