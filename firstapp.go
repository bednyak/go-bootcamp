package main

import (
	"fmt"
	"go-bootcamp/greeting"
)

//import "fmt"
//import "go-bootcamp/greeting"

func main() {
	//var greetingText string
	//greetingText = "Hello World!"

	//var greetingText string = "Hello World!"

	//greetingText := "Hello World!"
	luckyNumber := 17
	evenMoreLuckyNumber := luckyNumber + 5

	fmt.Println(greeting.GreetingText)
	fmt.Println(greeting.GreetingText)
	fmt.Println(luckyNumber)
	fmt.Println(evenMoreLuckyNumber)

	var newNumber float64 = float64(luckyNumber) / 3

	fmt.Println(newNumber)

	var defaultFloat float64 = 1.23456789123456789123456
	var smallFloat float32 = 1.23456789123456789123456

	fmt.Println(defaultFloat)
	fmt.Println(smallFloat)

	var firstRune rune = '&'
	fmt.Println(firstRune)
	fmt.Println(string(firstRune))

	var firstByte byte = 'a'
	fmt.Println(firstByte)

	firstName := "Andrii"
	lastName := "Biedniak"
	fullName := fmt.Sprintf("%v %v", firstName, lastName)
	age := 29

	fmt.Printf("Hi, I am %v and I am %v (Type: %T) years old.", fullName, age, age)

}
