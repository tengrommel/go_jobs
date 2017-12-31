package main

import "testing"

func add(x, y int) int {
	return x+y
}

func TestAdd(t *testing.T){
	r := add(2,4)
	if r!=6{
		t.Fatal("ok")
	}
	t.Log("test")
}
