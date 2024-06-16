package main
import "fmt"
type Stack struct{
	data []int 
} 
func (s *Stack)Pop()int {
	index:=len(s.data)-1
	item:=s.data[index]
	s.data=s.data[:index]

	return item
}

func (s *Stack)Push(item int){
	s.data=append(s.data,item)
	
}
func (s *Stack)IsEmpty()bool{
	return len(s.data)==0
}

func main() {
	stack:=Stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	for !stack.IsEmpty(){
		fmt.Println("poped :",stack.Pop())
	}
	
}