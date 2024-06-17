package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("enter the range :")
	fmt.Scan(&n)
	for i := 2; i < n; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	} else {
		for i := 2; i < n/2; i++ {
			if n%i == 0 {
				return false
			}
		}

	}
	return true
}
