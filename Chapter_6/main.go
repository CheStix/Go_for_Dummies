package main

import (
	"errors"
	"fmt"
	"strconv"
)

func insert(orig []int, index int, value int) ([]int, error) {
	if index < 0 {
		return nil, errors.New("Index cannot be less than 0")
	}
	if index > len(orig) {
		return append(orig, value), nil
	}
	orig = append(orig[:index+1], orig[index:]...)
	orig[index] = value
	return orig, nil
}

func delete(orig []int, index int) ([]int, error) {
	if index < 0 || index >= len(orig) {
		return nil, errors.New("Index out of range")
	}
	orig = append(orig[:index], orig[index+1:]...)
	return orig, nil
}

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

	s := make([]int, 2, 5)
	fmt.Println(s)
	fmt.Println(len(s), cap(s))

	t := []int{1, 2, 3, 4, 5}
	t = append(t, 6, 7, 8)
	fmt.Println(t)
	fmt.Println(len(t), cap(t))
	t = append(t, 9, 10)
	fmt.Println(t)
	fmt.Println(len(t), cap(t))
	u := t
	u[9] = 100
	t = append(t, 11)
	fmt.Println(t)
	fmt.Println(len(t), cap(t))
	fmt.Println(u)
	fmt.Println(len(u), cap(u))

	c := [3]string{"iOS", "Android", "Windows"}
	fmt.Println(c[0:2])
	fmt.Println(c[:2])
	fmt.Println(c[1:])
	fmt.Println(c[:])
	t = t[2:4]
	fmt.Println(t)
	fmt.Println(len(t), cap(t))
	t = t[1:3]
	fmt.Println(t)
	fmt.Println(len(t), cap(t))

	for i, v := range t {
		fmt.Println(i, v)

	}

	nums4 := nums1
	nums4[0] = 1
	fmt.Println(nums1, nums4)

	v := make([]int, len(t))
	copy(v, t)
	fmt.Println(t, v)

	v = make([]int, 1, 5)
	copy(v, t)
	fmt.Println(t, v)

	v = make([]int, 10)
	copy(v, t)
	fmt.Println(t, v)

	t = []int{1, 2, 3, 4, 5}
	t, err := insert(t, 2, 9)
	if err == nil {
		fmt.Println(t)
	} else {
		fmt.Println(err)
	}
	t, err = delete(t, 2)
	if err == nil {
		fmt.Println(t)
	} else {
		fmt.Println(err)
	}
}
