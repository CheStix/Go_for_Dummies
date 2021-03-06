package main

import (
	"fmt"
	"time"
)

func displayDate(format string, prefix string) {
	fmt.Println(prefix, time.Now().Format(format))
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func addNum(num1, num2 int) (sum int) {
	sum = num1 + num2
	return
}

func countOddEven(s string) (int, int) {
	odds, evens := 0, 0
	for _, c := range s {
		if int(c)%2 == 0 {
			evens++
		} else {
			odds++
		}
	}
	return odds, evens
}

func addNums(total int, nums ...int) int {
	fmt.Printf("%T ", nums)
	for _, n := range nums {
		total += n
	}
	return total
}

func fib() func() int {
	f1, f2 := 0, 1
	return func() int {
		f1, f2 = f2, f1+f2
		return f1
	}
}

func filter(arr []int, cond func(int) bool) []int {
	result := []int{}
	for _, v := range arr {
		if cond(v) {
			result = append(result, v)
		}
	}
	return result
}
func main() {
	displayDate("Mon 2006-01-02 15:04:05 MST", "Current Date and Time:")
	displayDate("15:04:05 Mon 02.01.2006 MST", "Current Time and Date:")

	x := 5
	y := 6
	swap(&x, &y)
	fmt.Println(x, y)

	s := addNum(x, y)
	fmt.Println(s)

	o, e := countOddEven("1234567")
	fmt.Println(o, e)

	fmt.Println(addNums(0, 1, 2, 3, 4, 5, 6))

	var i func() int
	i = func() int {
		return 5
	}
	fmt.Println(i())

	gen := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(gen())
	}

	a := []int{1, 2, 3, 4, 5}
	evens := filter(a, func(val int) bool {
		return val%2 == 0
	})

	fmt.Println(evens)
	evens = filter(a, func(val int) bool {
		return val > 3
	})
	fmt.Println(evens)
}
