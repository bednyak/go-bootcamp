package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	id          string
	title       string
	description string
	price       float64
}

func (prod *Product) store() {
	file, _ := os.Create(prod.id + ".txt")

	content := fmt.Sprintf(
		"ID: %v\nTitle: %v\nDescription: %v\nPrice: USD %.2f\n",
		prod.id,
		prod.title,
		prod.description,
		prod.price)

	file.WriteString(content)

	file.Close()
}

func (prod *Product) printData() {
	fmt.Printf("ID: %v\n", prod.id)
	fmt.Printf("Title: %v\n", prod.title)
	fmt.Printf("Description: %v\n", prod.description)
	fmt.Printf("Price: USD %.2f\n", prod.price)
}

func NewProduct(id string, t string, d string, p float64) *Product {
	return &Product{id, t, d, p}
}

func main() {
	firstProduct := Product{
		"first-product",
		"A Book",
		"A nice book",
		10.99,
	}

	secondProduct := NewProduct(
		"second-product",
		"A Carpet",
		"A beautiful carpet",
		99.99,
	)

	createdProduct := getProduct()

	createdProduct.printData()
	createdProduct.store()
	firstProduct.printData()
	secondProduct.printData()
}

func getProduct() *Product {
	fmt.Println("Please enter the product data.")
	fmt.Println("------------------------------")

	reader := bufio.NewReader(os.Stdin)

	idInput := readUserInput(reader, "Product ID: ")
	titleInput := readUserInput(reader, "Product Title: ")
	decriptionInput := readUserInput(reader, "Product Description: ")
	priceInput := readUserInput(reader, "Product Price: ")

	priceValue, _ := strconv.ParseFloat(priceInput, 64)

	product := NewProduct(idInput, titleInput, decriptionInput, priceValue)

	return product
}

func readUserInput(reader *bufio.Reader, promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)

	return userInput
}
