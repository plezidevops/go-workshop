package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("Hello World!")

	// out information to the user
	fmt.Println("BMI Calculator")
	fmt.Println("------------------------")

	// prompt for the user input (weight + height)
	fmt.Print("Enter your weight (lb): ")
	weightInput, _ := reader.ReadString('\n')

	fmt.Print("Enter height (feet): ")
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
