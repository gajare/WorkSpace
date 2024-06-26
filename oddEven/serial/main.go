package main

import (
	"fmt"
)

func main() {
	oddCh := make(chan bool)
	evenCh := make(chan bool)
	done := make(chan bool)

	go printOdd(oddCh, evenCh, done)
	go printEven(oddCh, evenCh, done)

	oddCh <- true // Start the sequence with the odd number

	<-done // Wait for the completion signal
}

func printOdd(oddCh, evenCh, done chan bool) {
	for i := 1; i <= 10; i += 2 {
		<-oddCh
		fmt.Println(i)
		evenCh <- true
	}
	done <- true
}

func printEven(oddCh, evenCh, done chan bool) {
	for i := 2; i <= 10; i += 2 {
		<-evenCh
		fmt.Println(i)
		if i >= 10 {
			done <- true
			return
		}
		oddCh <- true
	}
}
