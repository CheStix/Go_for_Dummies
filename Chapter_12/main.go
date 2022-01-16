package main

import (
	"fmt"
	"math/rand"
	"time"
)

//send data into channel
func sendData(ch chan string) {
	fmt.Println("Sending a string into channel...")
	//time.Sleep(2 * time.Second)
	ch <- "Hello"
	fmt.Println("String has been retrieved from channel..")
}

//getting data from the channel
func getData(ch chan string) {
	time.Sleep(2 * time.Second)
	fmt.Println("String retrieved from channel:", <-ch)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
	fmt.Println("Done and can continue to do other work")
}

func fib(n int, c chan int) {
	a, b := 1, 1
	for i := 0; i < n; i++ {
		c <- a
		a, b = b, a+b
		time.Sleep(2 * time.Second)
	}
	close(c)
}

func counter(n int, c chan int) {
	for i := 0; i < n; i++ {
		c <- i
		time.Sleep(time.Second)
	}
	close(c)
}
func main() {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	//fmt.Scanln()

	rand.Seed(time.Now().UnixNano())
	var s []int
	sliceSize := 10
	for i := 0; i < sliceSize; i++ {
		s = append(s, rand.Intn(100))
	}

	c := make(chan int, 5) //buffered channel of length 5
	partSize := 2
	parts := sliceSize / partSize
	i := 0
	for i < parts {
		go sum(s[i*partSize:(i+1)*partSize], c)
		i += 1
	}

	i = 0
	total := 0
	time.Sleep(time.Second)
	for i < parts {
		partialSum := <-c //read from channel
		fmt.Println("Partial sum:", partialSum)
		total += partialSum
		i += 1
	}
	fmt.Println("Total:", total)

	ch1 := make(chan int)
	ch2 := make(chan int)

	go fib(10, ch1)
	go counter(10, ch2)

	ch1Closed := false
	ch2Closed := false

	go func() {
		for {
			select {
			case n, ok := <-ch1:
				if !ok {
					//channel closed and drained
					ch1Closed = true
					if ch1Closed && ch2Closed {
						return
					}
				} else {
					fmt.Println("fib()", n)
				}
			case m, ok := <-ch2:
				if !ok {
					ch2Closed = true
					if ch1Closed && ch2Closed {
						return
					}
				} else {
					fmt.Println("counter()", m)
				}

			}
		}
	}()

	//for i := range ch1 {
	//	fmt.Println(i)
	//}
	//for i:=range ch2 {
	//	fmt.Println(i)
	//}
	fmt.Println("Do something")
	fmt.Scanln()
}
