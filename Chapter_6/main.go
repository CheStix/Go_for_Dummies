package main

import (
	"fmt"
	"strconv"
)

func main() {
	var nums1 [5]int // an array of int (5 elements)
	nums2 := [5]int{1, 2, 3, 4, 5}
	nums3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(nums1)
	fmt.Println(nums2)
	fmt.Println(len(nums3))

	var table [5][6]string
	for row := 0; row < 5; row++ {
		for column := 0; column < 6; column++ {
			table[row][column] = strconv.Itoa(row) + "," + strconv.Itoa(column)
		}
	}
	fmt.Println(table)

	var cube [4][3][3]string
	for row := 0; row < 4; row++ {
		for column := 0; column < 3; column++ {
			for depth := 0; depth < 3; depth++ {
				cube[row][column][depth] = strconv.Itoa(row) + strconv.Itoa(column) + strconv.Itoa(depth)
			}
		}
	}
	fmt.Println(cube)
}
