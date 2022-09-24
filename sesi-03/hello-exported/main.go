package main

import (
	"hello-exported/helpers"
)

func main() {
	helpers.Greet()

	var person = helpers.Person{}

	person.Invokegreet()
}
