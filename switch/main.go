package main

import "fmt"

func main() {
	fmt.Println("This is Switch Case:")
	// day := 5
	var day int
	fmt.Println("Enter any number:")
	error, _ := fmt.Scanf("%d", &day)
	// fmt.Scanf("%d", &day)
	if error != 0 {
		switch day {
		case 1:
			fmt.Println("Monday")
		case 2:
			fmt.Println("Tuesday")
		case 3:
			fmt.Println("Wednesday")
		case 4:
			fmt.Println("Thursday")
		case 5:
			fmt.Println("Friday")
		case 6:
			fmt.Println("Saturday")
		default:
			fmt.Println("Invalid day")
		}
	} else {
		fmt.Println("Error Handling")
	}
	// switch day {
	// case 1:
	// 	fmt.Println("Monday")
	// case 2:
	// 	fmt.Println("Tuesday")
	// case 3:
	// 	fmt.Println("Wednesday")
	// case 4:
	// 	fmt.Println("Thursday")
	// case 5:
	// 	fmt.Println("Friday")
	// case 6:
	// 	fmt.Println("Saturday")
	// default:
	// 	fmt.Println("Invalid day")
	// }

	// Consume the leftover newline character
	// _, _ = fmt.Scanf("%*c")

	// month := "January"
	month := "June"

	switch month {
	case "January", "February", "March":
		fmt.Println("Winter")
	case "April", "May", "June":
		fmt.Println("Spring")
	default:
		fmt.Println("Other Season")
	}

	temperature := 12
	switch {
	case temperature < 0:
		fmt.Println("Freezing")
	case temperature > 0 && temperature < 20:
		fmt.Println("Cold")
	case temperature > 20 && temperature < 25:
		fmt.Println("Warm")
	case temperature > 25:
		fmt.Println("HOT")

	}

}
