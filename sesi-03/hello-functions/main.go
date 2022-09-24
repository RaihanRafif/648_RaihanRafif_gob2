package main

import "fmt"

func main() {
	// greet("Airell", 23)
	// greet("airell", "jendereal sudirman")

	// var names = []string{"Airell", "Jordan"}
	// var printMsg = greet("Heii", names)
	// fmt.Println(printMsg)

	// var diameter float64 = 15
	// var area, circumference = calculate(diameter)

	// fmt.Println("Area:", area)
	// fmt.Println("Circumference:", circumference)

	// studentLists := print("Airell", "Nanda", "Maila", "Schannel", "Marco")
	// fmt.Printf("%v", studentLists)

	numberLists := []int{1, 2, 3, 4, 5, 6, 7, 8}

	result := sum(numberLists...)

	fmt.Println("result:", result)

	// profile("Airell", "pasta", "Ayam geprek", "ikan roa", "sate padang")
}

// func profile(name string, favFoods ...string) {
// 	mergeFavFoods := strings.Join(favFoods, " | ")

// 	fmt.Println("Hello there!! I'm", name)
// 	fmt.Println("I really love to eat", mergeFavFoods)
// }

func sum(numbers ...int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total
}

// func print(names ...string) []map[string]string {
// 	var result []map[string]string
// 	for i, v := range names {
// 		key := fmt.Sprintf("Student%d", i+1)
// 		temp := map[string]string{
// 			key: v,
// 		}
// 		result = append(result, temp)
// 	}
// 	return result
// }

// func calculate(d float64) (area float64, circumference float64) {
// 	area = math.Pi * math.Pow(d/2, 2)

// 	circumference = math.Pi * d
// 	return
// }

// func greet(msg string, names []string) string {
// 	var joinStr = strings.Join(names, " ")
// 	var result string = fmt.Sprintf("%s %s", msg, joinStr)

// 	return result
// }

// func greet(name, address string) {
// 	fmt.Println("hello there! My name is", name)
// 	fmt.Println("I live in", address)
// }

// func greet(name string, age int8) {
// 	fmt.Printf("Hello there! My name is %s and I'm %d years old", name, age)
// }
