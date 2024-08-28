package main

import "fmt"

func main() {
	//Array
	// arr := [4]int{1, 2, 3, 4}
	// fmt.Println("The elements in array:", arr)

	// //slice
	// slice := []int{10, 20, 30, 40}
	// fmt.Println("slice Initially:")
	// fmt.Println("The elements in slice:", slice)

	// //Adding an element to slice
	// slice = append(slice, 50, 60)
	// fmt.Println("Slice after adding some elements:")
	// fmt.Println("The elements in slice:", slice)

	// //creating slice from an array
	// sliceFromArray := arr[1:3]
	// fmt.Println("Slice created from Array")
	// fmt.Printf("The elements type in slice %T\n", sliceFromArray)
	// fmt.Println("The elements in slice:", sliceFromArray)

	numbers := make([]int, 3, 5)
	numbers = append(numbers, 34)
	numbers = append(numbers, 23)
	numbers = append(numbers, 12)
	fmt.Println("The elements in slice:", numbers)
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))

	stock := make([]int, 0)
	fmt.Println("Elements:", stock)
	fmt.Println("Length:", len(stock))
	fmt.Println("Capacity:", cap(stock))

}
