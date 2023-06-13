package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := 5
	b := 10
	c := 8

	a, b = generateRandomNumbers()

	sum := add(a, b)
	printMyNumber(c)

	fmt.Println("\nThe sum is", sum)
}

//func generateRandomNumbers() (int, int) {
//	randomNumber1 := rand.Intn(10)
//	randomNumber2 := rand.Intn(10)
//	return randomNumber1, randomNumber2
//}

//func add(num1 int, num2 int) int {
//	sum := num1 + num2
//	return sum
//}

func generateRandomNumbers() (r1 int, r2 int) {
	r1 = rand.Intn(10)
	r2 = rand.Intn(10)
	return
}

func add(num1 int, num2 int) (sum int) {
	sum = num1 + num2
	return
}

func printMyNumber(myNumber int) {
	fmt.Printf("My number is %v", myNumber)
}
