package main

import "fmt"

var greetingText = "Hello World!!!"

func main() {
	//var firstName string = "Andrii"
	//var lastName string = "Biedniak"
	firstName := "Andrii"
	lastName := "Biedniak"

	fmt.Println(firstName)
	fmt.Println(lastName)

	//var currentYear int = 2023
	//var birthYear int = 1994
	currentYear := 2023
	birthYear := 1994

	age := currentYear - birthYear

	fmt.Println(age)

	nextYear := currentYear + 1

	age = nextYear - birthYear

	fmt.Println(age)
}
