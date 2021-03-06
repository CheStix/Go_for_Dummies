package main

import (
	"fmt"
	"strings"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	max := 100
	a, b := 0, 1
	for b <= max {
		fmt.Println(b)
		a, b = b, a+b
	}

	for a, b := 0, 1; b <= max; a, b = b, a+b {
		fmt.Println(b)
	}

	for {
		fmt.Println("Enter QUIT to exit")
		var input string
		fmt.Print("Please enter a string:")
		fmt.Scanln(&input)
		if strings.ToUpper(input) == "QUIT" {
			break
		}
	}

	var OS [3]string
	OS[0] = "iOS"
	OS[1] = "Android"
	OS[2] = "Windows"
	for i, v := range OS {
		fmt.Println(i, v)
	}
	for _, v := range OS {
		fmt.Println(v)
	}
	for i, _ := range OS {
		fmt.Println(i)
	}
	for i := range OS {
		fmt.Println(i)
	}

	for pos, char := range "Hello, World!" {
		fmt.Println(pos, char)
		fmt.Printf("%d %c\n", pos, char)
	}

	for pos, char := range "こんにちは世界" {
		fmt.Printf("%c at byte location %d\n", char, pos)
	}

Outerloop:
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if i == 2 {
				continue Outerloop
			}
			if i == 4 {
				break Outerloop
			}
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
		fmt.Println("------------")
	}
}
