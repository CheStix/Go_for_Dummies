package main

import (
	"fmt"
	"unicode/utf8"
)

const publisher = "Wiley"

var (
	firstName string = "Wei-Meng"
	lastName  string = "Lee"
	age       int    = 25
)

func main() {
	var num1 = 5            //type inferred
	var num2 int            // explicitly typed
	var num3 float32        //floating point variable
	var name string         //string variable
	var raining bool        // boolean variable
	var rates float32 = 4.5 //declared as float and then initialized

	address := "The White House\n1600 Pennsylvania Avenue NW\nWashington, DC 20500"
	quotation := `"Any who never made
a mistake has never tried
anything new"
--Albert Einstein`
	str1 := "你好,世界"
	str2 := "こんにちは世界"
	str3 := "\u4f60\u597d\uff0c\u4e16\u754c"

	fmt.Println(num1)    // 5
	fmt.Println(num2)    // 0
	fmt.Println(num3)    // 0
	fmt.Println(name)    // "'
	fmt.Println(raining) // false
	fmt.Println(rates)   // 4.5
	fmt.Println(firstName, lastName, age)
	fmt.Println(publisher)
	fmt.Println(address)
	fmt.Println(quotation)
	fmt.Println(str1, len(str1), utf8.RuneCountInString(str1))
	fmt.Println(str1, len(str2), utf8.RuneCountInString(str2))
	fmt.Println(str1, len(str3), utf8.RuneCountInString(str3))
}
