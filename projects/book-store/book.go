package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

type Product struct {
	id string
	title string
	description string
	price float64
}

func (product Product) outputDetails() {
	fmt.Println("Production Information")
	fmt.Printf("ID: %v\nTitle: %v\nDescription: %v\nPrice: %.2f\n", 
	product.id, product.title, product.description, product.price)
}

func (store Product) saveToFile() {
	// get the current path

	// Create a file
	f, err := os.Create("product.txt")
	check(err)

	// save product to the file
	data := fmt.Sprintf("%v %v %v %0.2f", store.id, store.title, store.description, store.price)
	f.WriteString(data)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func NewProduct(id, title, description string, price float64) *Product {

	product := Product {
		id: id,
		title: title,
		description: description,
		price: price,
	}

	return &product
}

func main() {

	var product *Product

	id := getUserData("Enter the product ID: ")
	title := getUserData("Enter the product title: ")
	description := getUserData("Enter Product description: ")
	price := getUserData("Enter Product price: ")

	if p, err := strconv.ParseFloat(price, 64); err == nil {
		product = NewProduct(id, title, description,  p)
	} else {
		fmt.Println(err, p)
		fmt.Println("Invalid price")
	}
	
	product.outputDetails()
	product.saveToFile()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')

	var cleanedInput string;

	if runtime.GOOS == "windows" {
		cleanedInput = strings.Replace(userInput, "\r\n", "", -1)
	} else {
		cleanedInput = strings.Replace(userInput, "\n", "", -1)
	}
	return cleanedInput
}


// Your Tasks
// 1) Create a new type / data structure for storing product data (e.g. about a book)
//		The data structure should contain these fields:
//		- ID
//		- Title / Name
//		- Short description
//		- price (number without currency, we'll just assume USD)
// 2) Create concrete instances of this data type in two different ways:
//		- Directly inside of the main function
//		- Inside of main, by using a "creation helper function"
//		(use any values for title etc. of your choice)
//		Output (print) the created data structures in the command line (in the main function)
// 3) Add a "connected function" that outputs the data + call that function from inside "main"
// 4) Change the program to fetch user input values for the different data fields
//		and create only one concrete instance with that entered data.
//		Output that instance data (via the connected function) then.
// 5) Bonus: Add a connected "store" function that writes that data into a file
//		The filename should be the unique id, the function should be called at the
//		end of main.