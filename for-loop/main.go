package main

import "fmt"

func main() {
	// for i := 0; i <= 3; i++ {
	// 	fmt.Println("The value of i is", i)
	// }

	//Infinite loop with break statement ~ while loop
	// for{.....} creates an infinite loop
	/* counter := 0
	for {
		fmt.Println("Infinite Loop")
		counter++
		if counter == 3 {
			break
		}
	 }
	*/

	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("The index %d is and the value is %d\n", index, value)
	}
}
