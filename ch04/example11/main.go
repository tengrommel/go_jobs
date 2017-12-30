package main

import "fmt"

func testMap()  {
	//var a map[string]string = map[string]string{}
	a := make(map[string]string)
	a["abc"]  = "efg"
	a["abc"]  = "efd"
	a["abc1"]  = "efg"
	a["abc2"]  = "eg"
	a["ab2"]  = "eg"
	fmt.Println(a)
	fmt.Println(len(a))
}

func testMap2()  {
	a := make(map[string]map[string]string, 100) // 为了性能
	a["key1"] = make(map[string]string)
	a["key1"]["key2"]="value"
	fmt.Println(a)
}

func main() {
	testMap()
	testMap2()
}
