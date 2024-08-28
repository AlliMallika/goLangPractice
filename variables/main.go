package main

import (
	"MyLearning/myUtil"
	"fmt"
)

/*here import  "hello/myUtil"  tells Go to look for the myUtil package in the hello directory.
The PrintMessage function from the myUtil package is then used in the main function.

*/

func main() {
	fmt.Println("Hello from GoLang")
	myUtil.PrintMessage(("This is Mallika"))
	var name string = "Shiv"
	name = "Krsn"
	var version = "1.23.4"
	fmt.Println(name)
	fmt.Println(version)

	var isCheck bool = false
	fmt.Println(isCheck)
	isCheck = true
	fmt.Println("After reassignment of isCheck:", isCheck)

	var money int = 345200
	fmt.Println("Money in account is ", money)

	const pi = 23.12
	fmt.Println(pi)

	//short variable declaration
	Company := "IBM"
	fmt.Println(Company)

	//directly export to any other file and accessible in any file
	var Public = "data is important"

	//accessible within this file
	var private = "This is private variable"
	fmt.Println(Public)
	fmt.Println(private)
}
