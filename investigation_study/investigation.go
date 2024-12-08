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
    student = "Noman"
    age = 89
    isStudent = true
  )

  fmt.Println(student,age, isStudent)

  // so i cant use short hand notation in grouping in var or constant
  // const (
  //   CLOUDINARY_API_KEY := "9393939"
  //   CLOUD_NAME := "cloud"
  //   CLOUD_PASSWORD := "password"
  // )


  // if i want to use short hand notation with grouping i need to use this way without var or const keyword

    CLOUDINARY_API_KEY, CLOUD_NAME, CLOUD_PASSWORD := "9393939", "cloud", 909090990
    fmt.Println(CLOUDINARY_API_KEY, CLOUD_NAME, CLOUD_PASSWORD)
  
}