package main

import "fmt"

func main() {
	a := 0
	b := 1
	n := 10
	for i := 0; i < n; i++ {
		c := a + b
		fmt.Println(c)
		a = b
		b = c
	}
}
