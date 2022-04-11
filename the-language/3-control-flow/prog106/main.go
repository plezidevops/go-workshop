// break, goto, continue, fallthrough

package main

import "fmt"

func my_break() {
	for x := 1; x <= 6; x++ {
		fmt.Println("Number for:", x)

		if x == 4 {
			break
		}
	}
}

func my_goto() {
	var a int = 1

Label1:
	for a <= 5 {
		if a == 3 {
			a = a + 1
			goto Label1
		}
		fmt.Printf("Number goto is: %d\n", a)
		a++
	}
}

func my_continue() {
	var a int = 1
	for a <= 10 {
		if a == 7 {
			a++
			continue
		}
		fmt.Printf("Number continue is %d\n", a)
		a++
	}
}

func my_fallthrough() {
	switch number := 200; {
	case number < 300:
		fmt.Printf("%d is less than 300\n", number)
		fallthrough
	case number > 100:
		fmt.Printf("%d is greater than 100\n", number)
	}
}

func main() {
	my_break()
	my_goto()
	my_continue()
	my_fallthrough()
}
