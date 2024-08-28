package main

import "fmt"

// func divide(a, b float32) (float32, error) {
// 	if b != 0 {
// 		return a / b, nil
// 	}
// 	return 0, fmt.Errorf("denominator must not be zero")
// }

func divide(a, b float32) (float32, string) {
	if b != 0 {
		return a / b, "nil"
	}
	return 0, "here Denominator is Zero"
}

func main() {
	fmt.Println("Started error handling in functions")
	// ans, err := divide(10, 0)
	// if err != nil {
	// 	fmt.Println("Error Handling")
	// }

	ans, _ := divide(10, 0)
	fmt.Println("The division of two numbers:", ans)
}
