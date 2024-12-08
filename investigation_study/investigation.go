package main

import "fmt"

/*

1. First this i know about that i can't use short hand inference is grouping
2. Second thing i notice that if we create constants or normal var outside the main func compiler not complain me to use this variables

*/

const (
  USER_ID  string = "39299393"
  USER_EMAIL  string = "n@n.com"
  USER_PASSWORD  string = "nnnnnnnn"
)

func main() {

  // when ever i use var keyword to create variables go compiler complain me to use this variables
  var (
    student string = "Noman"
    age int = 89
    isStudent bool = true
  )

  // fmt.Println(student,age, isStudent)
}