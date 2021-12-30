package main

import "fmt"

type Node struct {
	value int
	next *Node
}

type LinkedList struct {
	len int
	head *Node
}

func (l *LinkedList) insertAtIndex(n *Node, index int) {
	if index == 0 {
		secondNode := l.head
		l.head = n
		n.next = secondNode
		l.len++
		return
	}

	currentIdx := 0
	currentNode := l.head

	for currentIdx != index - 1 {
		currentIdx++
		currentNode = currentNode.next
	}

	nextNode := currentNode.next
	currentNode.next = n
	n.next = nextNode
	l.len++
}

func (l *LinkedList) insertAtEnd(n *Node) {
	currentNode := l.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = n
	l.len++
}

func (l *LinkedList) readAtIndex(index int) (value int) {
	if index == 0 {
		return l.head.value
	}

	currentIdx := 0
	currentNode := l.head

	for currentIdx != index {
		currentNode = currentNode.next
		currentIdx++
	}
	return currentNode.value
}

func (l *LinkedList) deleteAtIndex(index int) (deleted *Node) {
	if index == 0 {
		deleted = l.head
		l.head = l.head.next
		return 
	}

	currentIdx := 0
	currentNode := l.head

	for currentIdx != index - 1 {
		currentNode = currentNode.next
		currentIdx++
	}

	deleted = currentNode.next
	currentNode.next = currentNode.next.next
	l.len--
	return
}

func (l *LinkedList) indexOf(n *Node) (index int) {
	currentNode := l.head

	for currentNode != n {
		if currentNode.next == nil {
			index = -1
			break
		}
		currentNode = currentNode.next
		index++
	}
	return
}

func (l *LinkedList) reverse() {
	var previousNode *Node
	currentNode := l.head
	
	for currentNode != nil {
		nextNode := currentNode.next
		currentNode.next = previousNode
		previousNode = currentNode
		currentNode = nextNode
	}

	l.head = previousNode
}

func (l *LinkedList) printAll() {
	if l.head == nil {
		fmt.Println("The list is empty!")
	}
	currentNode := l.head
	for currentNode != nil {
		fmt.Print(" -> ", currentNode.value)
		currentNode = currentNode.next
	}
	fmt.Println("")
}

func main() {
	node1 := &Node{value:1}
	node2 := &Node{value:2}
	node3 := &Node{value:3}
	node4 := &Node{value:4}
	node5 := &Node{value:5}
	node6 := &Node{value:6}

	list := &LinkedList{}
	list.printAll()
	list.insertAtIndex(node1, 0)

	fmt.Println("should be 1:", list.head.value)

	list.insertAtEnd(node2)
	list.insertAtEnd(node3)

	fmt.Println("should be 3:", list.head.next.next.value)
	list.insertAtIndex(node4, 1)
	fmt.Println("should be 4:", list.head.next.value)
	fmt.Println("should be 4:", list.readAtIndex(1))
	fmt.Println("should be 3:", list.readAtIndex(3))
	list.deleteAtIndex(3)
	fmt.Println("should be 2:", list.readAtIndex(2)) 
	list.insertAtEnd(node5)
	list.deleteAtIndex(2)
	fmt.Println("should be 5:", list.readAtIndex(2)) 

	fmt.Println("should be 0:", list.indexOf(node1))
	fmt.Println("should be 1:", list.indexOf(node4))
	fmt.Println("should be -1:", list.indexOf(node6))
	fmt.Print("should be -> 1 -> 4 -> 5:")
	list.printAll()
	list.reverse()
	fmt.Print("should be => 5 -> 4 -> 1:")
	list.printAll()
}