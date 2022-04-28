package main

import "fmt"

type Queue struct {
	data []int
}

func (q *Queue) Read() int {
	return q.data[0]
}

func (q *Queue) Push(value int) {
	q.data = append(q.data, value)
	return
}

func (q *Queue) Pop() int {
	removedElement := q.data[0]
	if len(q.data) > 1 {
		q.data = q.data[1:]
	} else {
		q.data = []int{}
	}
	return removedElement
}

func main() {
	newQ := Queue{data:[]int{}}
	newQ.Push(1)
	newQ.Push(2)
	newQ.Push(3)
	newQ.Push(4)
	fmt.Println(newQ) // 1 2 3 4
	newQ.Pop()
	fmt.Println(newQ) // 2 3 4
	newQ.Pop()
	fmt.Println(newQ) // 3 4
	fmt.Println(newQ.Read()) // 3
}