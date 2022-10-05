package main

import (
	"fmt"
)

func main() {
	c := make(chan string)

	// Tanda <- merupakan sebuah operator dari channel untuk proses pengiriman data dari Goroutine satu dengan yang lainnya.
	// c <- value

	// result := <-c

	// go introduce("Airell", c)

	// go introduce("Nanda", c)

	// go introduce("Mailo", c)

	// msg1 := <-c
	// fmt.Println(msg1)

	// msg2 := <-c
	// fmt.Println(msg2)

	// msg3 := <-c
	// fmt.Println(msg3)

	// ---------------------------------------------------
	// Channel with anonymous function
	students := []string{"Airell", "Mailo", "Indah"}

	// for _, v := range students {
	// 	go func(student string) {
	// 		fmt.Println("Student", student)
	// 		result := fmt.Sprintf("Hai, my name is %s", student)
	// 		c <- result
	// 	}(v)
	// }

	for _, v := range students {
		go introduce(v, c)
	}

	for i := 1; i <= 3; i++ {
		print(c)
	}

	close(c)
}

func print(c <-chan string) {
	fmt.Println(<-c)
}

// func print(c chan string) {
// 	fmt.Println(<-c)
// }

func introduce(student string, c chan<- string) {
	result := fmt.Sprintf("Hai, my name is %s", student)
	c <- result
}

// func introduce(student string, c chan string) {
// 	result := fmt.Sprintf("Hai, my name is %s", student)
// 	c <- result
// }
