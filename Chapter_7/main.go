package main

import (
	"fmt"
	"math"
)

type point struct {
	x float32
	y float32
	z float32
}

// a copy of the point struct is passed into the length() method
func (p point) length() float64 {
	return math.Sqrt(
		math.Pow(float64(p.x), 2) +
			math.Pow(float64(p.y), 2) +
			math.Pow(float64(p.z), 2))
}

// a reference to the struct is passed into it instead of a copy

func (p *point) move(deltax, deltay, deltaz float32) {
	p.x += deltax
	p.y += deltay
	p.z += deltaz
}
func newPoint(x, y, z float32) *point {
	p := point{x: x, y: y, z: z}
	return &p
}

func main() {
	var pt1 point
	pt1.x = 3.1
	pt1.y = 5.7
	pt1.z = 4.2
	fmt.Println(pt1.x, pt1.y, pt1.z)

	pt2 := point{
		x: 5.6,
		y: 3.8,
		z: 6.9,
	}
	pt3 := point{1.2, 2.3, 3.4}
	fmt.Println(pt2, pt3)

	pt4 := newPoint(7.8, 8.9, 9.0)
	fmt.Println(pt4)
	pt5 := pt4
	pt5.x = 0
	fmt.Println(pt4, pt5)
	pt6 := *pt4 // copy
	pt6.z = 0
	fmt.Println(pt4, pt6)
	pt7 := pt6  // copy
	pt8 := &pt7 //reference
	pt8.y = 0
	fmt.Println(pt8, pt7)

	fmt.Println(pt4.length())
	fmt.Println(pt4)
	pt4.move(0.1, 0.1, 0.1)
	fmt.Println(*pt4)
}
