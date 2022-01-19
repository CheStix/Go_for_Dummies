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

func main() {
	url := "http://data.fixer.io/api/latest?access_key=1b7c7ab128348573376f75d4e806f5c9"

	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			var result Result
			json.Unmarshal([]byte(body), &result)
			if result.Success {
				// create anrray to store all keys
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
}
