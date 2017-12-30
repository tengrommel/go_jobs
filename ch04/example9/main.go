package main

import "fmt"

func testSlice()  {
	var a [5]int = [...]int{1,2,3,4,5}
	s := a[1:]
	fmt.Printf("%p \n",&s)
	fmt.Printf("s=%p a[1]=%p",s, &a[1])
	fmt.Println(a)
	s = append(s, 10)
	s = append(s, 10)
	s = append(s, 10)
	s = append(s, 10)
	s = append(s, 10)
	s = append(s, 10)
	s = append(s, 10)
	s = append(s, 10)
	s = append(s, 10)
	s[1] = 10000
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Printf("%p \n",&s)
	fmt.Printf("s=%p a[1]=%p \n",s, &a[1])
	fmt.Println(a)
}

func main() {
	testSlice()
}
