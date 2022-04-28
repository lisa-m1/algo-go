package main

import "fmt"

type Stack struct {
	data []string
}

func (s *Stack) Pop() string {
	lastIdx := len(s.data) - 1
	removedElement := s.data[lastIdx]
	s.data = s.data[:lastIdx]
	return string(removedElement)
}

func (s *Stack) Push(c string) {
	s.data = append(s.data, c)
	return
}

func (s *Stack) Read() string {
	lastIdx := len(s.data) - 1
	return string(s.data[lastIdx])
}

func (s *Stack) Valid() bool {
	return len(s.data) != 0
}

func reverse(str string) (reversed string){
	myStack := Stack{[]string{}}
	for _, runeValue := range str {
		myStack.Push(string(runeValue))
	}

	for myStack.Valid() {
		reversed += myStack.Pop()
	}
	return
}

func main() {
	myString := "abcde"
	reversed := reverse(myString)
	fmt.Println(reversed)

}