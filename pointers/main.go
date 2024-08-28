package main

import "fmt"

//modify value by accessing the address
func modifyValueByReference(num *int) {
	*num = *num + 10
}

func main() {
	// var num int
	// num = 2
	// var ptr *int
	// ptr = &num

	// Initialize variable
	// testing git push
	num := 2

	// Create a pointer to the variable
	ptr := &num
	fmt.Println("The value in variable:", num)
	fmt.Println("Address in pointer:", ptr)
	// Print the value pointed to by the pointer (dereferencing)
	fmt.Println("Data through pointer:", *ptr)

	var pointer *int // default pointer is nil
	if pointer == nil {
		fmt.Println("Pointer is not assigned")
	}

	value := 20
	modifyValueByReference(&value)
	fmt.Println("Modified value:", value)
}
