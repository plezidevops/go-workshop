// handling error

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter your age: ")
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)
	userage, err := strconv.ParseInt(userInput, 0, 64)

	if err != nil {
		fmt.Println("Invalid input: ", err)
		return
	}

	fmt.Println(userage)
}
