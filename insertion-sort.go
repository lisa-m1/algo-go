package main

import "fmt"

func insertionSort(collection []int) []int {
	for idx := 1; idx < len(collection); idx++ { 
		emptySlot := idx 
		removedElement := collection[idx] 
		
		for currIdx := idx - 1; currIdx >= 0; currIdx-- {
			if removedElement > collection[currIdx] {
				collection[emptySlot] = removedElement
				break
			} else {
				collection[emptySlot] = collection[currIdx]
				emptySlot = currIdx
			}
		}
		if emptySlot == 0 {
			collection[emptySlot] = removedElement
		}
	}
	return collection
}


func main() {
	fmt.Println(insertionSort([]int{4, 6, 2, 1}))
	fmt.Println(insertionSort([]int{4, 6, 2, 1, 9, 7}))
	fmt.Println(insertionSort([]int{43, 36, 12, 61, 39, 75}))
	fmt.Println(insertionSort([]int{44, 16, 2, 31, 99, 7}))
}