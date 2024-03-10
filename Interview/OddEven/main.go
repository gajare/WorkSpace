package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Odd(ch chan bool) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		ch <- true
		fmt.Println(i)
		<-ch // Wait for the signal from the other goroutine before proceeding

	}
}

func Even(ch chan bool) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		ch <- true // Send signal (true) to allow the other goroutine to proceed
		fmt.Println(i)
		<-ch

	}
}
func main() {
	wg.Add(2)
	ch := make(chan bool, 5)
	go Even(ch)
	go Odd(ch)
	wg.Wait()
}