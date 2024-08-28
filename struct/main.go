package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	Email string
	Phone string
}
type Employee struct {
	EmployeeDetails        Person
	EmployeeContactDetails Contact
}

func main() {
	// var p Person
	// fmt.Println("Person", p)
	// p.FirstName = "Mallika"
	// p.LastName = "Alli"
	// p.Age = 23
	//	fmt.Println("******")
	//	fmt.Println("Person:", p)

	person2 := Person{
		FirstName: "Sravan",
		LastName:  "Alli",
		Age:       21,
	}
	fmt.Println("Second Person:", person2)

	//using new keyword
	person3 := new(Person)
	// var person3 = new(Person)
	person3.FirstName = "Shiv"
	person3.LastName = "Parvati"
	person3.Age = 35

	fmt.Println("Third Person:", person3)
	fmt.Println("Third Person FirstName:", person3.FirstName)
	fmt.Println("Third Person:", *person3)

	//Employee
	var employee1 Employee
	employee1.EmployeeDetails = Person{
		FirstName: "Mallika",
		LastName:  "Alli",
		Age:       23,
	}
	// employee1.EmployeeContactDetails = Contact{
	// 	Email: "allimallika19@gmail.com",
	// 	Phone: "8344842289",
	// }

	employee1.EmployeeContactDetails.Email = "allimallika19@gmail.com"
	employee1.EmployeeContactDetails.Phone = "8344842289"
	fmt.Println("Employee:", employee1)
}
