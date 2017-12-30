package main

import "fmt"

func bsort(a []int)  {
	for i:=0;i<len(a);i++ {
		for j:=1;j<len(a)-i;j++ {
			if a[j]>a[j-1] {
				a[j],a[j-1]= a[j-1], a[j]
			}
		}
	}
}


func main() {
	b := [10]int{8,7,5,4,2,19,10,9,3,11}
	bsort(b[:])
	fmt.Println(b)
}
