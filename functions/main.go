package main

import "fmt"

func testFunction() {
	fmt.Println("This is no arguments functions")
}
func add(a, b int) int {
	return a + b
}

// func add(a int, b int)(int){
// 	return a + b
// }

// func add(a int, b int)(result int)  {
// 	result = a + b
// 	return
// }

func main() {
	fmt.Println("Now we are learning about Functions in GoLang")
	testFunction()
	ans := add(3, 17)
	fmt.Println("The addition of two numbers:", ans)

}
