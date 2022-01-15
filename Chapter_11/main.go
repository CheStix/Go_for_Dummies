package main

import (
	"fmt"
	"time"
)

func say(s string, times int) {
	for i := 0; i < times; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i, s)
	}
}
func main() {
	go say("Hello", 3)
	go say("World", 2)
	fmt.Scanln()
}
