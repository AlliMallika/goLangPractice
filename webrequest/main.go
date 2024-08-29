package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Started learning WebRequest")

	//this file contains only GET HTTP Request
	res, err := http.Get("https://fakestoreapi.com/users")
	if err != nil {
		fmt.Println("Error in Response", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting Response:", res.Status)
		return
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error in Response Body:", err)
		return
	}
	fmt.Println("Response Body:", string(data))
}
