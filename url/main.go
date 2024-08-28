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
	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("RawQuery:", parsedURL.RawQuery)

	//modified URL components
	parsedURL.Path = "/productNames"
	parsedURL.RawQuery = "username='Mallika'&password='238@$*!meds'"

	//constructing a URL string from URL object
	newURL := parsedURL.String()
	fmt.Println("Modified URL:", newURL)
}
