package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("enter number :")
	fmt.Scanln(&n)
	fmt.Println("factorial of given number :", factorial(n))
}
func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return factorial(n-1) * n
	}
}
