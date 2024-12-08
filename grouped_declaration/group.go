package main

import "fmt"

var (
	isLoggedIn = true
	isAuthentic = false
	isAdmin = false
)

func main(){
	username, age := "Noman", 69
	fmt.Println(username, age)

	isAdmin = true

	fmt.Println(isAuthentic, isLoggedIn, isAdmin)

}