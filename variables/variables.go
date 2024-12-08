package main

import "fmt"

// There are different ways to create variables

// Fun fact I can create variables outside the main function ( entry point function )

var itWorks string = "WoW"
var doAgain = "Yo"
const itsPrettyGood string = "I love this language"
const awesome = "Go is awesome"

// So I cant declare shorthand variables outside the main function :(

// fistLanguage := "JavaScript"


func main(){
	fmt.Println(itWorks)
	fmt.Println(doAgain)
	fmt.Println(itsPrettyGood)
	fmt.Println(awesome)
	// fmt.Println(firstLanguage)

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

	// constants in Go

	const constantVariable string = "Some Secret"	
	fmt.Println(constantVariable)

	const JWT_SECRET_KEY string = "889283982982983983"
	fmt.Println(JWT_SECRET_KEY)


	// So i cant do this kind of crappy things

/*
	const assignMeLatter string
	assignMeLatter = "Some values" 
*/

}