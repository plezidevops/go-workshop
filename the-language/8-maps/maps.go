package main

import "fmt"

func main() {
	// websites := []string{"http://google.com", "http://ibm.com", "http://aws.com"}
	websites := map[string]string{
		"Google":              "https://google.com",
		"IBM":                 "https://ibm.com",
		"Amazon Web Services": "https://aws.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["Google"])

	// Adding key/value pair to the map
	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)

	// delete a key
	delete(websites, "Google")
	fmt.Println(websites)
}
