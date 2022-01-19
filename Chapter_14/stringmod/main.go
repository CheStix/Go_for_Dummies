package main

import (
	"fmt"
	qt "github.com/chestix/stringmod/quotes"
	str "github.com/chestix/stringmod/strings"
)

func main() {
	o, e := str.CountOddEven("12345")
	fmt.Println(o, e)

	fmt.Println(qt.GetEmoji("turtle"))
}
