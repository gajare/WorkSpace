package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	ch := make(chan bool, 5)
	go Odd(ch)
	go Even(ch)
	wg.Wait()

}

func Odd(ch chan bool) {
	defer wg.Done()
	for i := 0; i <= 10; i += 2 {
		
		ch <- true
		<-ch
		fmt.Printf("%d", i)
	}
}
func Even(ch chan bool) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		ch <- true
		<-ch
		fmt.Printf("%d", i)
	}
}
