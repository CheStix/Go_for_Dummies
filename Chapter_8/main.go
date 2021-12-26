package main

import (
	"fmt"
	"sort"
)

var heights map[string]int

func main() {
	heights = make(map[string]int)
	heights["Peter"] = 170
	heights["Joan"] = 168
	heights["Jan"] = 175
	fmt.Println(heights["Peter"], heights["Joan"], heights["Jan"])

	heights := map[string]int{
		"Peter": 170,
		"Joan":  168,
		"Jan":   175,
	}
	fmt.Println(heights)
	fmt.Println(heights["Pit"])
	if v, ok := heights["Pit"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key does not exist")
	}

	if _, ok := heights["Jan"]; ok {
		delete(heights, "Jane")
	} else {
		fmt.Println("Key does not exist")
	}
	fmt.Println(len(heights))

	for k, v := range heights {
		fmt.Println(k, v)
	}

	// get all keys in map
	var keys []string
	for k := range heights {
		keys = append(keys, k)
	}
	fmt.Println(keys)

	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, heights[k])
	}
}
