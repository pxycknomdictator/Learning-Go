package main

import "fmt"

// There are different ways to create variables

func main(){
	// variable declaration + explicit type declaration
	var username string = "Noman"
	fmt.Println(username)

	var age int = 69
	fmt.Println(age)


	// variable declaration but type infer by go
	var isDev = true
	fmt.Println(isDev)

	var language = "Go"
	fmt.Println(language)

	// variable decleration + type declaration in short hand
	developer := "Web-dev"
	fmt.Println(developer)

	student := true
	fmt.Println(student)
}