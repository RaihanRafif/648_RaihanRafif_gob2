package main

import (
	"fmt"
)

func main() {
	// var currentYear = 2001
	// if age := currentYear - 1998; age < 17 {
	// 	fmt.Printf("kamu belum bile membuat kartu sim => %d", age)
	// } else {
	// 	fmt.Println("kamu sudah boleh membuat kartu sim")
	// }

	var score = 8
	// switch score {
	// case 8:
	// 	fmt.Println("Perfect")
	// case 7:
	// 	fmt.Println("awesome")
	// default:
	// 	fmt.Println("not bad")
	// }

	// switch score {
	// case score == 8:
	// 	fmt.Println("Perfect")
	// case (score < 8) && (score > 3):
	// 	fmt.Println("awesome")
	// 	fallthrough
	// case score < 5:
	// 	fmt.Println("It is ok, but please study harder")
	// default:
	// 	{
	// 		fmt.Println("study harder")
	// 		fmt.Println("You don't have a good score yet")
	// 	}
	// }

	// switch {
	// case score == 8:
	// 	fmt.Println("perfect")
	// case (score < 8) && (score > 3):
	// 	fmt.Println("not bad")
	// 	fallthrough
	// case score < 5:
	// 	fmt.Println("It is ok, but please study harder")
	// default:
	// 	{
	// 		fmt.Println("study harder")
	// 		fmt.Println("You don't have a good score yet")
	// 	}
	// }

	if score > 7 {
		switch score {
		case 10:
			fmt.Println("perfect")
		default:
			fmt.Println("nice!")
		}
	} else {
		if score == 5 {
			fmt.Println("not bad")
		} else if score == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if score == 0 {
				fmt.Println("Try Harder!")
			}
		}
	}

}
