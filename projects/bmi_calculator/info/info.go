package info

import "fmt"

const MainTitle = "BMI Calculator"
const Separator = "------------------------"
const WeightInputPrompt = "Enter your weight (lb): "
const HeightInputPrompt = "Enter height (feet): "

// display message
func DisplayUserInfo() {
	fmt.Println(MainTitle)
	fmt.Println(Separator)
}
