package main

import (
	"fmt"
)

func main() {
	// city := "jakarta"

	// for i := 0; i < len(city); i++ {
	// 	fmt.Printf("index: %d, byte:%d\n", i, city[i])
	// }

	// var city []byte = []byte{74, 97, 107, 97, 114, 116, 97}
	// fmt.Println(string(city))

	// str1 := "makan"
	// str2 := "mânca"

	// fmt.Printf("str1 byte length => %d\n", len(str1))
	// fmt.Printf("str2 byte length => %d\n", len(str2))

	// fmt.Printf("str1 byte length => %d\n", utf8.RuneCountInString(str1))
	// fmt.Printf("str2 byte length => %d\n", utf8.RuneCountInString(str2))

	str := "mânca"
	for i, s := range str {
		fmt.Printf("index => %d,string => %s\n", i, string(s))
	}
}
