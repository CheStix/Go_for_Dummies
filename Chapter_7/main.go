package main

import "fmt"

type point struct {
	x float32
	y float32
	z float32
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
}
