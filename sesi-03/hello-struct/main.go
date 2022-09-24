package main

import "fmt"

// type Employee struct {
// 	name     string
// 	age      int
// 	division string
// }

// type Employee struct {
// 	division string
// 	person   Person
// }

type Person struct {
	name string
	age  int
}

func main() {
	// var employee Employee

	// employee.name = "Airell"
	// employee.age = 23
	// employee.division = "Curriculum developer"

	// fmt.Println(employee.name)
	// fmt.Println(employee.age)
	// fmt.Println(employee.division)

	// -----------------------------------------------------------------------------

	// var employee1 = Employee{}
	// employee1.name = "Airell"
	// employee1.age = 23
	// employee1.division = "Curriculum developer"

	// var employee2 = Employee{name: "Ananda", age: 23, division: "Finance"}

	// fmt.Printf("Employee1: %+v\n", employee1)
	// fmt.Printf("Employee2: %+v\n", employee2)

	// -----------------------------------------------------------------------------
	// var employee1 = Employee{name: "Ananda", age: 23, division: "Finance"}
	// var employee2 *Employee = &employee1

	// fmt.Println("Employee1 name:", employee1.name)
	// fmt.Println("Employee2 name:", employee2.name)

	// fmt.Println(strings.Repeat("#", 20))

	// employee2.name = "Ananda"

	// fmt.Println("Employee1 name:", employee1.name)
	// fmt.Println("Employee2 name:", employee2.name)

	// -----------------------------------------------------------------------------

	// var employee1 = Employee{}

	// employee1.person.name = "Airell"
	// employee1.person.age = 23
	// employee1.division = "Curriculum developer"

	// fmt.Printf("%+v", employee1)

	// -----------------------------------------------------------------------------

	// var employee1 = struct {
	// 	person   Person
	// 	division String
	// }{}
	// employee1.person = Person{name: "Airell", age: 23}
	// employee1.division = "web dev"

	// var employee1 = struct {
	// 	person   Person
	// 	division String
	// }{
	// 	person:   Person{name: "slene", age: 23},
	// 	division: "web dev",
	// }

	// fmt.Printf("Employee1: %+v\n", employee1)
	// fmt.Printf("Employee2: %+v\n", employee2)

	// -----------------------------------------------------------------------------
	// var people = []Person{
	// 	{name: "asas", age: 23},
	// 	{name: "zxcxzc", age: 23},
	// 	{name: "qwwe", age: 23},
	// }

	// for _, v := range people {
	// 	fmt.Printf("%+v\n", v)
	// }

	// -----------------------------------------------------------------------------
	var employee = []struct {
		person   Person
		division string
	}{
		{person: Person{name: "asdasd", age: 23}, division: "web dev"},
		{person: Person{name: "qweqwe", age: 23}, division: "golang dev"},
		{person: Person{name: "xzczxc", age: 23}, division: "kotlin dev"},
	}

	for _, v := range employee {
		fmt.Printf("%+v\n", v)
	}
}
