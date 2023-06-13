package main

import (
	"fmt"
)

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	var myHobbies [3]string = [3]string{"playing-computer-games", "read-books", "outdoor-walking"}

	fmt.Println(myHobbies)

	fmt.Println(myHobbies[0])

	fmt.Println(myHobbies[1:3])

	var mainHobbies1 = myHobbies[0:2]

	var mainHobbies2 = myHobbies[:2]

	fmt.Println(mainHobbies1)

	fmt.Println(mainHobbies2)

	fmt.Println(mainHobbies2[1:3])

	var myGoals []string = []string{"learn-go", "have-fun"}

	fmt.Println(myGoals)

	myGoals[1] = "have-knowledge"

	fmt.Println(myGoals)

	var updatedMyGoals []string = append(myGoals, "be-good")

	fmt.Println(updatedMyGoals)

	var p1 = Product{
		"id1",
		"Table",
		100.33,
	}

	var p2 = Product{
		"id2",
		"Chair",
		40.83,
	}

	var p3 = Product{
		"id3",
		"TV",
		999.93,
	}
	var productList []Product = []Product{p1, p2}

	fmt.Println(productList)

	productList = append(productList, p3)

	fmt.Println(productList)
}
