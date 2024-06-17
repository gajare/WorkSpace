package main

import "fmt"

func main() {
	s1 := []int{1, 2, 34, 56, 4, 1, 2, 5, 6, 7, 8, 9, 4, 5, 6, 1, 2}
	fmt.Println("s1 :", s1)
	Rpeated(s1)
}
func Rpeated(s1 []int) {
	m1 := make(map[int]int)
	for _, v := range s1 {
		m1[v]++
	}
	s2 := []int{}
	for k, v := range m1 {
		if v > 1 {
			s2 = append(s2, k)
		}
	}
	fmt.Println("s2 :", s2)
}
