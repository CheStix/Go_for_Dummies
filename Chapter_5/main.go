package main

import (
	"fmt"
	"time"
)

func displayDate(format string, prefix string) {
	fmt.Println(prefix, time.Now().Format(format))
}
func main() {
	displayDate("Mon 2006-01-02 15:04:05 MST", "Current Date and Time:")
	displayDate("15:04:05 Mon 02.01.2006 MST", "Current Time and Date:")
}