package main

import (
	"bmi/info"
)

func main() {

	// out information to the user
	info.DisplayUserInfo()

	// prompt for the user input (weight + height)
	weightInput, heightInput := getUserInput()

	// calculate bmi
	bmi := calculateBMI(weightInput, heightInput)

	// Output the calculated BMI
	printBMI(bmi)
}
