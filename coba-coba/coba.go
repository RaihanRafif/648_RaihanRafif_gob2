package main

func main() {
	TheFunction(2)
}

func TheFunction(a int) bool {
	if a%2 == 0 {
		return true
	} else {
		return false
	}
}
