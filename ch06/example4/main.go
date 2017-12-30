package main

import (
	"reflect"
	"fmt"
)

func test(b interface{})  {
	t := reflect.TypeOf(b)
	fmt.Println(t)

	v := reflect.ValueOf(b)
	fmt.Println(v)
	k := v.Kind()
	fmt.Println(k)

	iv := v.Interface()
	fmt.Println(iv)
}

func main() {
	var a int = 200
	test(a)

}
