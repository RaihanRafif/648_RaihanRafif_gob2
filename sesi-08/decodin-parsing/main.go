package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func main() {
	var jsonString = `[
	{
		"full_name" : "raihan rafif",
		"email" : "raihan@gmail.com",
		"age" : 23
	},
	{
		"full_name" : "zee jkt48",
		"email" : "jkt48@gmail.com",
		"age" : 22
	}
	]
	`
	// var result Employee
	// var result = temp.(map[string]interface{})
	var result []Employee
	// var temp interface{}

	var err = json.Unmarshal([]byte(jsonString), &result)
	// var err = json.Unmarshal([]byte(jsonString), &temp)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// fmt.Println("full_name:", result.FullName)
	// fmt.Println("email:", result.Email)
	// fmt.Println("age:", result.Age)

	// fmt.Println("full_name:", result["full_name"])
	// fmt.Println("email:", result["email"])
	// fmt.Println("age:", result["age"])
	for i, v := range result {
		fmt.Printf("Index %d: %+v\n", i+1, v)
	}
}
