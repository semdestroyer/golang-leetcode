package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
}
func singleNumber(nums []int) int {
	digitsMap := make(map[int][]int)
	keys := make([]int, 0)
	for _, dig := range nums {
		digitsMap[dig] = append(digitsMap[dig], dig)
	}
	for k := range digitsMap {
		keys = append(keys, k)
	}
	for _, k := range keys {
		if len(digitsMap[k]) == 1 {
			return digitsMap[k][0]
		}
	}

	return -1

}

//todo: оптимизировать
