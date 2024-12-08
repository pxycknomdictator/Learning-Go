package main

import "fmt"

var (
	isLoggedIn bool = true
	isAuthentic bool = false
	isAdmin bool = false
)

const (
	JWT_SECRET_KEY string = "Some secret key"
	MONGODB_URL string = "mongodb://127.0.0.1:27017"
	API_KEY string = "Some Api key"
)

func main(){
	username, age := "Noman", 69
	fmt.Println(username, age)

	isAdmin = true
	fmt.Println(isAuthentic, isLoggedIn, isAdmin)

	
	fmt.Println(JWT_SECRET_KEY, MONGODB_URL, API_KEY)

}