package main

import (
	"fmt"
)

func main() {
	// var fruits = []string{"apple", "banana", "mango"}
	// _ = fruits

	// var fruits = make([]string, 3)
	// _ = fruits

	// fruits[0] = "apple"
	// fruits[1] = "banana"
	// fruits[2] = "mango"

	// fruits = append(fruits, "apple", "banana", "mango")

	// fmt.Printf("%#v", fruits)

	// var fruits1 = []string{"apple", "banana", "mango"}
	// var fruits2 = []string{"druian", "pineapple", "starfruit"}

	// fruits1 = append(fruits1, fruits2...)

	// nn := copy(fruits1, fruits2)

	// fmt.Println("fruits1 =>", fruits1)
	// fmt.Println("fruits2 =>", fruits2)
	// fmt.Println("Copied elements =>", nn)

	// fmt.Printf("%#v", fruits1)

	// var fruits1 = []string{"apple", "banana", "mango", "pineapple"}

	// var fruits2 = fruits1[1:4]
	// fmt.Printf("%#v\n", fruits2)

	// var fruits3 = fruits1[0:]
	// fmt.Printf("%#v\n", fruits3)

	// var fruits4 = fruits1[:3]
	// fmt.Printf("%#v\n", fruits4)

	// var fruits5 = fruits1[:]
	// fmt.Printf("%#v\n", fruits5)

	// fruits1 = append(fruits1[:3], "rambutan")
	// fmt.Printf("%#v", fruits1)

	// var fruits1 = []string{"apple", "mango", "durian", "banana", "starfruit"}

	// var fruits2 = fruits1[2:4]

	// fruits2[0] = "rambutan"
	// fmt.Println("fruits1 =>", fruits1)
	// fmt.Println("fruits2 =>", fruits2)

	// fmt.Println("Fruits cap:", cap(fruits1))
	// fmt.Println("Fruits len:", len(fruits1))

	// fmt.Println(strings.Repeat("#", 20))

	// var fruits2 = fruits1[0:3]

	// fmt.Println("Fruits cap:", cap(fruits2))
	// fmt.Println("Fruits len:", len(fruits2))

	// fmt.Println(strings.Repeat("#", 20))

	// var fruits3 = fruits1[1:]
	// fmt.Println("Fruits cap:", cap(fruits3))
	// fmt.Println("Fruits len:", len(fruits3))
	cars := []string{"ford", "honda", "audi", "Rango Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"
	fmt.Println("cars", cars)
	fmt.Println("newCars", newCars)
}
