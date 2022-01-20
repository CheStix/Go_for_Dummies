package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

type Result struct {
	Success   bool
	Timestamp int
	Base      string
	Date      string
	Rates     map[string]float64
}

type Error struct {
	Success bool
	Error   struct {
		Code int
		Type string
		Info string
	}
}

var apis map[int]string

// channel to store map[int]interface{}
var c chan map[int]interface{}

func fetchData(API int) {
	url := apis[API]
	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err == nil {

			var result map[string]interface{}
			json.Unmarshal([]byte(body), &result)
			var re = make(map[int]interface{})

			switch API {
			case 1: // for the Fixer API
				if result["success"] == true {
					re[API] = result["rates"].(map[string]interface{})["USD"]
				} else {
					re[API] = result["error"].(map[string]interface{})["info"]
				}
				// store the result into the channel
				c <- re
				fmt.Println("Result for API 1 stored")
			case 2: // for the openweathermap.org API
				if result["main"] != nil {
					re[API] = result["main"].(map[string]interface{})["temp"]
				} else {
					re[API] = result["message"]
				}
				// store the result into the channel
				c <- re
				fmt.Println("Result for API 2 stored")
			}
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
}

func main() {
	url := "http://data.fixer.io/api/latest?access_key=1b7c7ab128348573376f75d4e806f5c9"

	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			var result Result
			json.Unmarshal([]byte(body), &result)
			if result.Success {
				// create array to store all keys
				keys := make([]string, 0, len(result.Rates))
				// get all the keys
				for k := range result.Rates {
					keys = append(keys, k)
				}
				// sort the keys
				sort.Strings(keys)
				// print the keys and values in alphabetical order
				for _, k := range keys {
					fmt.Println(k, result.Rates[k])
				}
				/*
					for i, v := range result.Rates {
						fmt.Println(i, v)
					}
				*/
			} else {
				var err Error
				json.Unmarshal([]byte(body), &err)
				fmt.Println(err.Error.Info)
			}
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
	fmt.Println("Done")
	// creates a channel to store the results from the API calls
	c = make(chan map[int]interface{})
	apis = make(map[int]string)
	apis[1] = "http://data.fixer.io/api/latest?access_key=" + "1b7c7ab128348573376f75d4e806f5c9"
	apis[2] = "http://api.openweathermap.org/data/2.5/weather?" + "q=Moscow,ru&units=metric" + "&appid=c9ac1919d90a6331c4a6b7fcccb38184"

	go fetchData(1)
	go fetchData(2)

	// we expect two results in the channel
	for i := 0; i < 2; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Done!")
	fmt.Scanln()
}
