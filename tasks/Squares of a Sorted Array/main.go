package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
}

func sortedSquares(nums []int) []int {
	res := make([]int, 0)
	for _, num := range nums {
		res = append(res, num*num)
	}
	sort.Ints(res)

	return res
}

//TODO: переработать под два указателя
