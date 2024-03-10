package main

import "fmt"

func main() {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	target := 1

	index, found := LineraSearch(arr, target)
	if found {
		fmt.Println("target Found in the array at :", index)
	} else {
		fmt.Println("target NOT Found in the array")
	}
}
func LineraSearch(arr []int, target int) (int, bool) {
	for i, value := range arr {
		if value == target {
			return i, true
		}
	}
	return -1, false
}
