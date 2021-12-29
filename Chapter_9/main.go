package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Firstname string
	Lastname  string
	Details   struct {
		Height int
		Weight float32
	}
}

type Rates struct {
	Base   string `json:"base currency"`
	Symbol string `json:"destination currency"`
}

func main() {
	var person People
	jsonString1 := `{"firstname":"Wei-Meng",
					"lastname":"Lee"}`
	err := json.Unmarshal([]byte(jsonString1), &person)
	if err == nil {
		fmt.Println(person.Firstname)
		fmt.Println(person.Lastname)
	} else {
		fmt.Println(err)
	}

	var persons []People
	jsonString2 := `[{"firstname":"Wei-Meng", "lastname":"Lee", "details":{"height":175, "weight":70.0}},
					{"firstname":"Mickey", "lastname":"Mouse", "details":{"height":105, "weight":85.5}}]`
	json.Unmarshal([]byte(jsonString2), &persons)
	for _, person := range persons {
		fmt.Println(person.Firstname)
		fmt.Println(person.Lastname)
		fmt.Println(person.Details.Height)
		fmt.Println(person.Details.Weight)
	}

	jsonString3 :=
		`{
			"base currency":"EUR",
 			"destination currency":"USD"
		  }`
	var rates Rates
	json.Unmarshal([]byte(jsonString3), &rates)
	fmt.Println(rates.Base)
	fmt.Println(rates.Symbol)

	jsonString4 :=
		`{
		 "success": true,
		 "timestamp": 1588779306,
		 "base": "EUR",
		 "date": "2020-05-06",
		 "rates": {
			 "AUD": 1.683349,
			 "CAD": 1.528643,
			 "GBP": 0.874757,
			 "SGD": 1.534513,
			 "USD": 1.080054
			 }
		 }`
	var result map[string]interface{}
	json.Unmarshal([]byte(jsonString4), &result)
	fmt.Println(result["success"])
	rates1 := result["rates"]
	fmt.Println(rates1) // map[AUD:1.683349 CAD:1.528643 GBP:0.874757 SGD:1.534513 USD:1.080054]
	currencies := rates1.(map[string]interface{})
	SGD := currencies["SGD"]
	fmt.Println(SGD)
}
