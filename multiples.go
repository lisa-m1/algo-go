// https://projecteuler.net/problem=1
// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. 
// The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.

package main

import "fmt"

func findMultiplesSum(num int) (sum int) {
	for i := 1; i < num; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			sum += i
		}
	}
	return
}

func findMultiplesSumv2(num int) (sum int) {
	for i := 3; i < num; i += 3 {
		sum += i
	}

	for i := 5; i < num; i += 5 {
		if i % 3 != 0 {
			sum += i
		}
	}
	return
}

func main() {
	fmt.Println(findMultiplesSum(1000)) // 233168
	fmt.Println(findMultiplesSumv2(1000)) // 233168
}