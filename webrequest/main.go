package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Started learning WebRequest")

	res, err := http.Get("https://fakestoreapi.com/users")
	if err != nil {
		fmt.Println("Error in Response", err)
		return
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error in Response Body:", err)
		return
	}
	fmt.Println("Response Body:", string(data))
}
