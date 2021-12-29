package main

import (
	"fmt"
	"sort"
)

type kv struct {
	key   string
	value int
}

type kvPairs []kv

var heights map[string]int

func (p kvPairs) Len() int {
	// returns the length of the collection
	return len(p)
}

func (p kvPairs) Less(i, j int) bool {
	// indicate the first value (height) must be smaller than the second value
	return p[i].value < p[j].value
}

func (p kvPairs) Swap(i, j int) {
	// swap the items in the collection
	p[i], p[j] = p[j], p[i]
}

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

	p := make(kvPairs, len(heights))
	i := 0
	for k, v := range heights {
		p[i] = kv{k, v}
		i++
	}
	fmt.Println(p)
	sort.Sort(p)
	for _, v := range p {
		fmt.Println(v)
	}
}
