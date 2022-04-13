package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

const mainTitle = "BMI Calculator"
const separator = "------------------------"
const weightInputPrompt = "Enter your weight (lb): "
const heightInputPrompt = "Enter height (feet): "

func main() {

	// out information to the user
	fmt.Println(mainTitle)
	fmt.Println(separator)

	// prompt for the user input (weight + height)
	fmt.Print(weightInputPrompt)
	weightInput, _ := reader.ReadString('\n')

	fmt.Print(heightInputPrompt)
	heightInput, _ := reader.ReadString('\n')

	// save the user input in variables
	weightInput = strings.TrimSuffix(weightInput, "\n")
	heightInput = strings.TrimSuffix(heightInput, "\n")

	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)

	// calculate the BMI (weight / (height * height))
	bmi := weight / (height * height)

	// Output the calculated BMI
	fmt.Printf("You BMI: %.2f", bmi)
}
