package main

import (
	"fmt"
)

func main() {

	ch := make(chan bool, 5)
	go Odd(ch)
	go Even(ch)

}

func Odd(ch chan bool) {

	for i := 0; i <= 10; i += 2 {

		ch <- true
		<-ch
		fmt.Printf("%d", i)
	}
}
func Even(ch chan bool) {
	for i := 1; i <= 10; i += 2 {
		ch <- true
		<-ch
		fmt.Printf("%d", i)
	}
}
