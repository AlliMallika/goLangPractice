package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Now learning about URL")
	myURL := "https://fakestoreapi.com/products?limit=5"
	fmt.Printf("Type of URL %T\n", myURL)

	parsedURL, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("can't parse URL", err)
		return
	}
	fmt.Println("ParsedURL:", parsedURL)
	fmt.Printf("Type of ParsedURL:%T\n", parsedURL)

}
