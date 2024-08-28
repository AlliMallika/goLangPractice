package main

import "fmt"

func main() {
	age := 23
	name := "Mallika"
	fmt.Println("Name:", name, "Age:", age)
	fmt.Println("This is Println function")

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Println("__________")
	fmt.Printf("Name: %s, Age: %d", name, age)
}
