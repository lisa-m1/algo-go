// https://projecteuler.net/problem=3
// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

package main

import "fmt"

func isPrime(num int) bool {
	if num == 1 {
		return false
	}

	for i := 2; i < num; i++ {
		if num % i == 0 {
			return false
		}
	}
	
	return true
}

func primeFactor(num int) (prime int) {
	i := 2

	for num != 1 {
		isMultiple := false

		for num % i == 0 {
			num = num / i
			isMultiple = true
		}

		if isMultiple && isPrime(i) {
			prime = i
		}

		i++
	}

	return
}

func main() {
	fmt.Println(primeFactor(600851475143)) // 6857
}