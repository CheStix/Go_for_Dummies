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

type dob struct {
	day   int
	month int
	year  int
}

type people struct {
	name  string
	email string
	dob   dob
}

var members map[int]people
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

	members = make(map[int]people)
	members[1] = people{
		name:  "Mary Smith",
		email: "marysmith@mail.com",
		dob: dob{
			day:   17,
			month: 3,
			year:  1990,
		},
	}
	members[2] = people{
		name:  "John Smith",
		email: "johnsmith@mail.com",
		dob: dob{
			day:   9,
			month: 12,
			year:  1988,
		},
	}
	members[3] = people{
		name:  "Janet Doe",
		email: "janetdoe@mail.com",
		dob: dob{
			day:   1,
			month: 12,
			year:  1988,
		},
	}
	members[4] = people{
		name:  "Adam Jones",
		email: "adamjones@mail.com",
		dob: dob{
			day:   19,
			month: 8,
			year:  2001,
		},
	}
	for k, v := range members {
		fmt.Println(k, v.name, v.email, v.dob)
	}
	// get all keys in members
	var keys2 []int
	for k := range members {
		keys2 = append(keys2, k)
	}
	// sort the keys
	sort.Ints(keys2)
	// copy the values in members to the slice
	var sliceMembers []people
	for _, k := range keys2 {
		sliceMembers = append(sliceMembers, members[k])
	}
	fmt.Println(sliceMembers)

	sort.SliceStable(sliceMembers, func(i, j int) bool {
		// compare year
		if sliceMembers[i].dob.year != sliceMembers[j].dob.year {
			return sliceMembers[i].dob.year < sliceMembers[j].dob.year
		}
		// compare month
		if sliceMembers[i].dob.month != sliceMembers[j].dob.month {
			return sliceMembers[i].dob.month < sliceMembers[j].dob.month
		}
		return sliceMembers[i].dob.day < sliceMembers[j].dob.day
	})

	for _, v := range sliceMembers {
		fmt.Println(v.name, v.email, v.dob)
	}
}
