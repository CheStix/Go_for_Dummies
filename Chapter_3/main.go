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

	dayOfWeek := ""
	switch num {
	case 1:
		dayOfWeek = "Monday"
	case 2:
		dayOfWeek = "Tuesday"
	case 3:
		dayOfWeek = "Wednesday"
	case 4:
		dayOfWeek = "Thursday"
	case 5:
		dayOfWeek = "Friday"
	case 6:
		dayOfWeek = "Saturday"
	case 7:
		dayOfWeek = "Sunday"
	default:
		dayOfWeek = "--error--"
	}
	fmt.Println(dayOfWeek)

	score := 79
	grade := ""
	switch {
	case score < 50:
		grade = "F"
	case score < 60:
		grade = "D"
	case score < 70:
		grade = "C"
	case score < 80:
		grade = "B"
	default:
		grade = "A"
	}

	switch grade {
	case "A":
		fallthrough
	case "B":
		fallthrough
	case "C":
		fallthrough
	case "D":
		fmt.Println("Passed")
	case "F":
		fmt.Println("Failed")
	default:
		fmt.Println("Absent")
	}

	switch grade {
	case "A", "B", "C", "D":
		fmt.Println("Passed")
	case "F":
		fmt.Println("Failed")
	default:
		fmt.Println("Undefined")
	}

}
