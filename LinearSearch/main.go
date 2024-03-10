package main

import "fmt"

func main() {
	arr:=[]int{1,6,2,8,3,6,4,5,9,10}
	target:=7
	res:=LineraSearch(arr,target)
	if res!=0{
		fmt.Println("target Found in the array")
	}else{
		fmt.Println("target NOT Found in the array")
	}
}
func LineraSearch(arr[]int,target  int)int{
	for v:=range arr{
		if v==target{
			return v
		}
	}
	return 0
}
