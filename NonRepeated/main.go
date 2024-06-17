package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 1, 4, 4, 5, 5, 4, 6, 7, 8, 9}
	fmt.Println("s1 :", s1)
	NonRepeated(s1)
}
func NonRepeated(s1 []int) {
	m1 := make(map[int]int)
	for _, value := range s1 {
		m1[value]++
	}
	//key is more important
	for k, v := range m1 {
		if v == 1 {
			fmt.Println("k :", k)
		}

	}
}
