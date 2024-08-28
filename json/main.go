package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	fmt.Println("Now we are learning JSON")
	person := Person{Name: "Mallika", Age: 23}
	fmt.Println("Person Data is:", person)

	//convert person into JSON Encoding (Marshalling)
	jsonData, error := json.Marshal(person)
	if error != nil {
		fmt.Println("Erron in JSON Encoding", error)
		return
	}
	fmt.Println("Person in JSON FormData:", string(jsonData))

	//decoding (Unmarshalling)
	var personData Person
	error = json.Unmarshal(jsonData, &personData)
	if error != nil {
		fmt.Println("Error in Unmarshalling:", error)
		return
	}
	fmt.Println("Person data in Unmarshalling:", personData)
}
