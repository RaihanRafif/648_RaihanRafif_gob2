package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (r rectangle) perimeter() float64 {
	return 2 * (math.Pi + r.width)
}

func print(t string, s shape) {
	fmt.Printf("%s area %v\n", t, s.area())
	fmt.Printf("%s perimeter %v\n", t, s.perimeter())
}

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func main() {
	var c1 shape = circle{radius: 5}
	// var r1 shape = rectangle{width: 3, height: 2}

	// c1.(circle).volume()

	// fmt.Printf("Type of c1: %T", c1)
	// fmt.Printf("Type of r1: %T", r1)

	// fmt.Println("Circle area", c1.area())
	// fmt.Println("Circle perimeter", c1.perimeter())

	// fmt.Println("Circle area", r1.area())
	// fmt.Println("Circle perimeter", r1.perimeter())

	// print("Rectangle", c1)
	// print("Circle", r1)

	// --------------------------
	value, ok := c1.(circle)

	if ok == true {
		fmt.Printf("Circle value: %+v\n", value)
		fmt.Printf("Circle volume: %f\n", value.volume())
	}

}
