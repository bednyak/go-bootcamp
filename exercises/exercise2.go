package main

import "fmt"

func main() {
	const PI = 3.14
	circleRadius := 5

	var circleCircumference float64 = 2 * PI * float64(circleRadius)

	fmt.Println(circleCircumference)
	fmt.Printf("For a radius of %v, the circle circumference is %.2f.", circleRadius, circleCircumference)
}
