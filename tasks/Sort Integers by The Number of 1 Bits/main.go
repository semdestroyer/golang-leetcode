package main

import (
	"fmt"
	"math/bits"
	"sort"
)

func main() {
	fmt.Println(sortByBits([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(sortByBits([]int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}))
}

func sortByBits(arr []int) []int {
	res := make([]int, 0)
	keys := make([]int, 0)
	digMap := make(map[int][]int)

	for _, digit := range arr {
		digMap[bits.OnesCount(uint(digit))] = append(digMap[bits.OnesCount(uint(digit))], digit)
	}

	for key := range digMap {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	for _, key := range keys {
		sort.Ints(digMap[key])
		res = append(res, digMap[key]...)
	}

	return res
}

//TODO: очень говнянное и неоптимизированное решение, но  сегодня я себя плохо чувствую и позволю набыдлокодить