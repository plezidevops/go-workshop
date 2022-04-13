package main

import (
	"bmi/info"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// obtain user input
func getUserInput() (weightInput string, heightInput string) {

	fmt.Print(info.WeightInputPrompt)
	weightText, _ := reader.ReadString('\n')

	fmt.Print(info.WeightInputPrompt)
	heightText, _ := reader.ReadString('\n')

	// save the user input in variables
	weightInput = strings.TrimSuffix(weightText, "\n")
	heightInput = strings.TrimSuffix(heightText, "\n")

	return
}

// calculate BMI
func calculateBMI(weightInput, heightInput string) float64 {

	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)

	// calculate the BMI (weight / (height * height))
	return weight / (height * height)
}

func printBMI(bmi float64) {
	fmt.Printf("You BMI: %.2f", bmi)
}
