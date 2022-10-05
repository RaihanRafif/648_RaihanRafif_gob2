package main

import (
	"fmt"
	"os"
)

func main() {
	// defer fmt.Println("defer function starts to execute")
	// fmt.Println("hai eveyone")
	// fmt.Println("Welcome back to Go learning center")

	// callDeferFunc()
	// fmt.Println("Hai everyone!!")
	defer fmt.Println("Invoke with defer")
	fmt.Println("Before exiting")
	os.Exit(1)
}

func callDeferFunc() {
	defer deferFunc()
}

func deferFunc() {
	fmt.Println("Defer func starts to execute")
}
