package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
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

	strings.Replace(userInput, "\n", "", -1)

	if userInput == "1" ||
		userInput == "2" ||
		userInput == "3" ||
		userInput == "4" {
		return userInput, nil
	} else {
		return "", errors.New("INVALID INPUT")
	}
}

func calculateSumUpToNumber() {}

func calculateFactorial() {}

func calculateSumManually() {}

func calculateListSum() {}

func main() {
	selectChoice, err := getUserChoice()

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
