package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func say(s string, times int) {
	for i := 0; i < times; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i, s)
	}
}

var balance int64
var mutex = &sync.Mutex{}

func credit() {
	for i := 0; i < 10; i++ {
		atomic.AddInt64(&balance, 100)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func debit() {
	for i := 0; i < 5; i++ {
		atomic.AddInt64(&balance, -100)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func main() {
	//go say("Hello", 3)
	//go say("World", 2)

	balance = 200
	fmt.Println("Initial balance is", balance)
	go credit()
	go debit()
	fmt.Scanln()
	fmt.Println(balance)
}
