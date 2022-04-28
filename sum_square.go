// https://projecteuler.net/problem=6

// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

package main

import "fmt"

func sumSquare(num int) int {
	var sumSq int
	var squareSum int
	
	for i := 1; i <= num; i++ {
		sumSq += i * i
		squareSum += i
	}

	return (squareSum * squareSum) - sumSq
}

func main() {
	fmt.Println(sumSquare(100)) // 25164150
}