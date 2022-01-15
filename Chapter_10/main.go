package main

import (
	"fmt"
	"math"
)

type DigitsCounter interface {
	CountOddEven() (int, int)
}

type DigitString string

func (ds DigitString) CountOddEven() (int, int) {
	odds, evens := 0, 0
	for _, digit := range ds {
		if digit%2 == 0 {
			evens++
		} else {
			odds++
		}
	}
	return odds, evens
}

type Circle struct {
	radius float64
	name   string
}
type Square struct {
	length float64
	name   string
}
type Triangle struct {
	base   float64
	height float64
}
type Shape interface {
	Area() float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
func (c Circle) Circumference() float64 {
	return 2 * math.Pi * c.radius
}
func (s Square) Area() float64 {
	return math.Pow(s.length, 2)
}
func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func calculateArea(listOfShapes []Shape) {
	for _, shape := range listOfShapes {
		fmt.Println("Area of shape is ", shape.Area())
	}
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%v %v (%d years old)", p.FirstName, p.LastName, p.Age)
}
func main() {
	s := DigitString("123456789")
	fmt.Println(s.CountOddEven())

	var d DigitsCounter
	d = s
	fmt.Println(d.CountOddEven())

	c1 := Circle{radius: 5, name: "c1"}
	s1 := Square{length: 6, name: "s1"}
	t1 := Triangle{base: 6, height: 8}

	shapes := []Shape{c1, s1, t1}
	calculateArea(shapes)

	me := Person{"Wei-meng", "Lee", 34}
	fmt.Println(me)
}
