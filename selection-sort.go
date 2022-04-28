package main

import "fmt"

func selectionSort(collection []int) []int {
	for i, num := range collection { 
		min := num	
		var minIdx = i 
		for idx, number := range collection[(i + 1):] { 
			if number < min {
				min = number
				minIdx = idx + i + 1
			}
		}
		
		if min != num {
			collection[i], collection[minIdx] = min, num
		}
	}
	return collection
}

func main() {
	fmt.Println(selectionSort([]int{4, 6, 2, 1, 9, 7}))
	fmt.Println(selectionSort([]int{43, 36, 12, 61, 39, 75}))
	fmt.Println(selectionSort([]int{44, 16, 2, 31, 99, 7}))
}