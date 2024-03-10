package main

import "fmt"

func main() {
	arr := []int{1, 2, 4, 8, 16, 32, 64, 128, 256}
	target := 13
	index := binarySearch(arr, target)
	if index != -1 {
		fmt.Println("Element found at index :", index)
	} else {
		fmt.Println("Element Not found")
	}
}
func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == target {
			return mid // Target found at index mid
		} else if arr[mid] < target {
			low = mid + 1 // Search in the right half
		} else {
			high = mid - 1 // Search in the left half
		}
	}

	return -1 // Target not found
}
