package main

import (
	"fmt"
	"strconv"
	"strings"

	"bmi/info"
)

func main() {

	// out information to the user
	fmt.Println(info.MainTitle)
	fmt.Println(info.Separator)

	// prompt for the user input (weight + height)
	fmt.Print(info.WeightInputPrompt)
	weightInput, _ := reader.ReadString('\n')

	fmt.Print(info.HeightInputPrompt)
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
