// Return the content type of an http header

package main

import (
	"fmt"
	"net/http"
)

func main() {
	ctype, err := contentType("https://google.com")

	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(ctype)
	}
}

// return the value of Content-Type header return by making an HTTP GET request to url
func contentType(url string) (string, error) {

	// make an http call
	res, err := http.Get(url)

	// check for errors
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	// get content type
	// ctype := strings.Join(res.Header["Content-Type"], ";")
	ctype := res.Header.Get("Content-Type")

	// make sure content type is not empty
	if ctype == "" {
		return "", fmt.Errorf("can't find Content-Type header")
	}

	// return
	return ctype, err

}
