package main

import "fmt"

func raining() bool {
	fmt.Println("Check if it is raining now...")
	return true
}

func snowing() bool {
	fmt.Println("Check if it is snowing now...")
	return true
}
func main() {
	num := 6
	condition := num%2 == 1
	if condition {
		fmt.Println("Number is odd")
	} else {
		fmt.Println("Number is even")
	}

	if raining() || snowing() {
		fmt.Println("Stay indoors!")
	}

}
