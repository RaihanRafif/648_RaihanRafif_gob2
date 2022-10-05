package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var number = 23
	// var reflectValue = reflect.ValueOf(number)

	// fmt.Println("Tipe variabel : ", reflectValue.Type())
	// fmt.Println("nilai variabel : ", reflectValue.Interface())

	// if reflectValue.Kind() == reflect.Int {
	// 	fmt.Println("Nilai variabel : ", reflectValue.Int())
	// }

	// var nilai = reflectValue.Interface().(int)

	// ---------------------------------------------------
	var s1 = &student{Name: "Raihan Rafif", Grade: 2}
	fmt.Println("nama :", s1.Name)

	var reflectValue = reflect.ValueOf(s1)
	var method = reflectValue.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("Wick"),
	})
	fmt.Println("nama:", s1.Name)
}

func (s *student) SetName(name string) {
	s.Name = name
}
