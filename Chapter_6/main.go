package main

import "fmt"

func main() {
	var nums1 [5]int // an array of int (5 elements)
	nums2 := [5]int{1, 2, 3, 4, 5}
	nums3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(nums1)
	fmt.Println(nums2)
	fmt.Println(len(nums3))
}
