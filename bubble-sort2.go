package main

import "fmt"

func bubbleSort(collection []int) []int {
	var sorted bool
	length := len(collection) - 1

	for !sorted {
		sorted = true
		for i := 0; i < length; i++ {
			if collection[i] > collection[i + 1] {
				collection[i], collection[i + 1] = collection[i + 1], collection[i]
				sorted = false
			}
		}
		length--
	}
	return collection
}

func main() {
	fmt.Println(bubbleSort([]int{4, 3, 6, 1, 9}))
	fmt.Println(bubbleSort([]int{44, 23, 46, 431, 19}))
	fmt.Println(bubbleSort([]int{144, 223, 46, 31, 9}))
}
