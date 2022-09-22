package main

import "fmt"

func main() {
	// var person map[string]string

	// person = map[string]string{}

	// person["name"] = "Airell"
	// person["age"] = "23"
	// person["address"] = "jalan sudirman"

	// fmt.Println("name:", person["name"])
	// fmt.Println("age", person["age"])
	// fmt.Println("address", person["address"])

	// var person = map[string]string{
	// 	"name":    "airell",
	// 	"age":     "23",
	// 	"address": "Jalan sudirman",
	// }

	// fmt.Println("name:", person["name"])
	// fmt.Println("age", person["age"])
	// fmt.Println("address", person["address"])

	// for key, value := range person {
	// 	fmt.Println(key, ":", value)
	// }

	// fmt.Println("Before deleting", person)
	// delete(person, "age")
	// fmt.Println("after deleting", person)

	// value, exist := person["age"]

	// if exist {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("value does'nt exist")
	// }

	// delete(person, "age")

	// value, exist = person["age"]

	// if exist {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("value has been deleted")
	// }

	var people = []map[string]string{
		{"name": "airell", "age": "23"},
		{"name": "nanda", "age": "23"},
		{"name": "mailo", "age": "15"},
	}

	for i, person := range people {
		fmt.Printf("Index: %d, name: %s, age %s\n", i, person["name"], person["age"])
	}
}
