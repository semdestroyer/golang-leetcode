package main

import "fmt"

func main() {
	fmt.Println(countBits(2))
	fmt.Println(countBits(5))
}

func countBits(n int) []int {
	var bits = make([]int, 0)
	for i := 0; i < n+1; i++ {
		bits = append(bits, countOnes(i))
	}
	return bits
}

func countOnes(n int) int {
	count := 0
	for n > 1 {
		if n%2 == 1 {
			count++
		}
		n = n / 2
	}
	if n%2 == 1 {
		count++
	}
	return count
}
