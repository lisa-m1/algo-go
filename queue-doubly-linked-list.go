package main

import "fmt"

type Node struct {
	previousLink *Node
	data int
	nextLink *Node
}

type Queue struct {
	firstNode *Node
	lastNode *Node
}

func (q *Queue) queue(n *Node) {
	if q.lastNode == nil {
		q.firstNode = n
		q.lastNode = n
	} else {
		previous := q.lastNode
		q.lastNode.nextLink = n
		q.lastNode = n
		q.lastNode.previousLink = previous
	}
}

func (q *Queue) read() (int) {
	return q.firstNode.data
}

func (q *Queue) dequeue() (*Node) {
	first := q.firstNode
	q.firstNode = q.firstNode.nextLink
	q.firstNode.previousLink = nil
	return first
}

func main() {
	node := &Node{data: 1}
	node2 := &Node{data: 2}
	node3 := &Node{data: 3}
 node4 := &Node{data: 4}

	newQueue := &Queue{}

	newQueue.queue(node2)
	newQueue.queue(node)
	newQueue.queue(node3)
	newQueue.queue(node4)

	fmt.Println(newQueue.read()) // 2
	fmt.Println(node2.nextLink) // node
	fmt.Println(node.nextLink) // node3
	fmt.Println(node3.nextLink) // node4
	fmt.Println(node4.nextLink) // nil
	fmt.Println(node2.previousLink) // nil
	fmt.Println(node.previousLink) // node2
	fmt.Println(node3.previousLink) // node
	fmt.Println(node4.previousLink) // node3
	fmt.Println(newQueue.dequeue()) // node2
	fmt.Println(newQueue.read()) // 1
	fmt.Println(newQueue.dequeue()) // node
	fmt.Println(newQueue.read()) // 3
	fmt.Println(newQueue.dequeue()) // node3
	fmt.Println(newQueue.read()) // 4
}