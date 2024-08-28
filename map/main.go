package main

import "fmt"

func main() {
	studentGrades := make(map[string]int)
	studentGrades["Shiv"] = 90
	studentGrades["Krsn"] = 80

	fmt.Println("Key-value pairs in map:", studentGrades)
	fmt.Println("Marks of Shiv:", studentGrades["Shiv"])
	fmt.Println("Marks of Krsn:", studentGrades["Krsn"])
	studentGrades["Krsn"] = 85
	fmt.Println("New Marks of Krsn:", studentGrades["Krsn"])

	delete(studentGrades, "Krsn")
	fmt.Println("New Marks of Krsn:", studentGrades["Krsn"])
	fmt.Println("Marks of Mallika:", studentGrades["Mallika"])

	fmt.Println("Checking if a key exists or not")
	//checking if a key exists
	grades, exists := studentGrades["Mallika"]
	fmt.Println("Grade of Mallika:", grades)
	fmt.Println("Mallika Exists or not:", exists)

	for index, value := range studentGrades {
		fmt.Printf("Key of student %s and the marks of a student %d\n", index, value)
	}
}
