package main

import (
	"net/http"
	"fmt"
)

func Hello(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("handle hello")
	fmt.Fprintf(w, "hello")
}
func login(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("handle login")
	fmt.Fprintf(w, "login")
}

func main() {
	http.HandleFunc("/", Hello)
 	http.HandleFunc("/user/login", login)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err!= nil{
		fmt.Println("err:", err)
	}
}
