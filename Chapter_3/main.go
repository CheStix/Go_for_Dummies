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

func doSomething() (int, bool) {
	//...
	//just an example of some return values
	return 5, false
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

	if v, err := doSomething(); !err {
		fmt.Println(v)
	} else {
		fmt.Println("handle the error")
	}
}
