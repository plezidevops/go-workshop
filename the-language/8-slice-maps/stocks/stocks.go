package main

import "fmt"

func main() {

	stocks := map[string]float64{
		"AMZN": 2087.98,
		"GOOG": 2540.85,
		"MSFT": 287.70,
	}

	fmt.Println(len(stocks))

	fmt.Println(stocks["GOOG"])

	stocks["MBN"] = 234

	value, ok := stocks["NFT"]

	if !ok {
		fmt.Println(value)
	} else {
		println(ok)
	}

	for value, i := range stocks {
		fmt.Printf("%v --> %v\n", value, i)
	}
}
