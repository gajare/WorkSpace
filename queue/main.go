package main

import "fmt"
type Queue struct {
	data []int
}

func (q *Queue) Add(item int) {
	q.data = append(q.data, item)
}
func (q *Queue) Delete() int {
	item := q.data[0]
	q.data = q.data[1:]
	return item
}

func (q *Queue) Dispaly() {
	for _, v := range q.data {
		fmt.Print(v)
	}
}
func main() {
	queue := Queue{}
	queue.Add(11)
	queue.Add(22)
	queue.Add(33)
	queue.Dispaly()
	fmt.Println("\n", queue.Delete())
	queue.Dispaly()
	fmt.Println("\n", queue.Delete())
	queue.Dispaly()

}
