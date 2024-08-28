package main

import (
	"fmt"
)

func main() {
	fmt.Println("Now started if-else condition in GoLang")
	x := 20
	if x > 20 {
		fmt.Println("x is greater than 20")
	} else {
		fmt.Println("x is less than or equal to 20")
	}

	var z int
	fmt.Println("Enter any number:")
	_, err := fmt.Scanf("%d", &z)
	if err != nil {
		fmt.Println("Reading input", err)
	}

	if z > 10 {
		fmt.Println("z is greater than 10")
	} else if z > 5 && z <= 10 {
		fmt.Println("z is greater than 5 and less than or equal to 10")
	} else {
		fmt.Println("z is less than or  equal to 5")
	}
}
