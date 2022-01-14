package main

import "fmt"

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

func main() {
	s := DigitString("123456789")
	fmt.Println(s.CountOddEven())

	var d DigitsCounter
	d = s
	fmt.Println(d.CountOddEven())
}
