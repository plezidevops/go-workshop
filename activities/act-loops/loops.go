// The code needs to refactor

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getUserChoice() (string, error) {
	fmt.Println("PLease enter your choice")
	fmt.Println("1) And up all the numbers of to number X")
	fmt.Println("2) Build the factorial up to number X")
	fmt.Println("3) Sum up manually entered numbers")
	fmt.Println("4) Sum up a list of entered numbers ")

	userInput, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	userInput = strings.Replace(userInput, "\n", "", -1)

	if userInput == "1" ||
		userInput == "2" ||
		userInput == "3" ||
		userInput == "4" {
		return userInput, nil
	} else {
		return "", errors.New("INVALID INPUT")
	}
}

func calculateSumUpToNumber() {
	fmt.Print("Please enter the numbers: ")
	chosenNumber, err := getInputNumber()

	if err != nil {
		fmt.Println("Invalid Input")
		return
	}

	sum := 0
	for i := 1; i <= chosenNumber; i++ {
		sum += i
	}

	fmt.Printf("Result: %v", sum)
}

func calculateFactorial() {

	fmt.Print("Please enter the numbers: ")
	chosenNumber, err := getInputNumber()

	if err != nil {
		fmt.Println("Invalid Input")
		return
	}

	sum := 1

	for i := 1; i <= chosenNumber; i++ {
		sum *= i
	}

	fmt.Printf("Result: %v", sum)
}

func calculateSumManually() {

	fmt.Print("Please enter another number (enter 0 or anything else to exit): ")
	enteredNumber, err := getInputNumber()

	if err != nil {
		fmt.Println("Invalid Input")
		return
	}

	result := 0

	for enteredNumber > 0 {

		result += enteredNumber
		fmt.Print("Please enter another number (enter 0 or anything else to exit): ")
		enteredNumber, err = getInputNumber()

		if err != nil {
			fmt.Println("Invalid Input")
			return
		}
	}
	fmt.Printf("Result: %v", result)
}

func calculateListSum() {

	numbers := getNumberList()
	sum := 0

	for n := range numbers {
		num, _ := strconv.ParseInt(numbers[n], 0, 64)
		sum += int(num)
	}
	fmt.Println(sum)
}

func getNumberList() (numbers []string) {

	fmt.Print("Please enter the numbers: ")
	inputNumbers, _ := reader.ReadString('\n')

	inputNumbers = strings.Replace(inputNumbers, "\n", "", -1)
	numbers = strings.Split(inputNumbers, " ")

	return
}

func getInputNumber() (int, error) {
	inputNumber, err := reader.ReadString('\n')

	if err != nil {
		return 0, err
	}

	inputNumber = strings.Replace(inputNumber, "\n", "", -1)
	chosenNumber, err := strconv.ParseInt(inputNumber, 0, 64)

	if err != nil {
		return 0, nil
	}

	return int(chosenNumber), err
}

func main() {
	selectChoice, err := getUserChoice()

	fmt.Println(err)

	if err != nil {
		fmt.Println("Invalid choice, exiting!")
		return
	}

	switch selectChoice {

	case "1":
		calculateSumUpToNumber()
	case "2":
		calculateFactorial()
	case "3":
		calculateSumManually()
	case "4":
		calculateListSum()
	default:
		fmt.Println("Something not right!")
	}
}
