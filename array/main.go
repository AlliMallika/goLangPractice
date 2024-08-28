package main

import "fmt"

func main() {

	// var names [2]string
	// names[0] = "Shiv"
	// names[1] = "Krsn"
	// fmt.Println("Names in array:", names)

	numbers := [4]int{23, 56, 24, 13} // Initialize and assign values to an array in single line
	fmt.Println("The numbers in array are:", numbers)
	fmt.Println("The element in index 2:", numbers[2])
	fmt.Println("The length of the array:", len(numbers))
}
