package main

import "fmt"

func main() {

	// 1) Create a new array (!) that contains three hobbies you have
	// Output (print) that array in the command line.
	myHobbies := [3]string{"model air plane", "go programming", "robotics"}
	fmt.Println(myHobbies)

	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	fmt.Println(myHobbies[0])
	fmt.Println(myHobbies[:2])
	fmt.Println(myHobbies[1:])

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	hobbiesSlice := myHobbies[:2]
	fmt.Println(hobbiesSlice)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	hobbiesSlice = myHobbies[1:]
	fmt.Println(hobbiesSlice)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	goal := []string{"play piano", "biking"}
	fmt.Println(goal)

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	goal[1] = "stock trade"
	goal = append(goal, "golf")
	fmt.Println(goal)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	type Product struct {
		title string
		id    string
		price float64
	}

	product := []Product{
		{
			title: "keyboard",
			id:    "ID101",
			price: 15.95,
		},
		{
			title: "monitor",
			id:    "ID102",
			price: 9.99,
		},
	}

	// prod := new(Product)
	// prod.title = "latop"
	// prod.id = "ID103"
	// prod.price = 899.99

	prod := Product{
		title: "laptop",
		id:    "ID103",
		price: 999.99,
	}

	newProduct := append(product, prod)

	for _, p := range newProduct {
		fmt.Println(p)
	}

}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
