package main

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

type Book struct {
	ID               string  `json:"id"`
	Title            string  `json:"title"`
	ShortDescription string  `json:"shortDescription"`
	Price            float64 `json:"price"`
}

func (book *Book) outputBookInfo() {
	fmt.Printf("ID: %v | Title: %v | Description: %v | Price: %v$\n", book.ID, book.Title, book.ShortDescription, book.Price)
}

func main() {

	var bookId string = generateUniqueId()
	var bookTitle string = "Title1"
	var bookShortDescription string = "Some awesome description"
	var bookPrice float64 = 222.99

	var newBook Book
	var newBook2 Book

	newBook = Book{
		ID:               bookId,
		Title:            bookTitle,
		ShortDescription: bookShortDescription,
		Price:            bookPrice,
	}

	var productId2 = getUserData("Please enter product ID: ")
	var productTitle2 = getUserData("Please enter product title: ")
	var productShortDescription2 = getUserData("Please enter product short description: ")
	var productPrice2 = getUserData("Please enter product price: ")

	var price, _ = strconv.ParseFloat(productPrice2, 64)
	newBook2 = *creationHelperFunction(productId2, productTitle2, productShortDescription2, price)

	fmt.Printf("ID: %v | Title: %v | Description: %v | Price: %v$\n", newBook.ID, newBook.Title, newBook.ShortDescription, newBook.Price)
	newBook2.outputBookInfo()
	newBook2.storeBook()
}

func creationHelperFunction(id string, title string, shortDesc string, price float64) *Book {
	product := Book{
		id,
		title,
		shortDesc,
		price,
	}

	return &product
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')

	cleanedInput := strings.Replace(userInput, "\n", "", -1)

	return cleanedInput
}

func (book *Book) storeBook() error {

	bookData, err := json.Marshal(book)
	if err != nil {
		panic(err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	filePath := filepath.Join(cwd, book.ID)

	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write(bookData)
	if err != nil {
		panic(err)
	}

	return nil
}

func generateUniqueId() string {
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	id := hex.EncodeToString(randomBytes)
	return id
}
