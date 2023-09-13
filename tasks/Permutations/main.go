package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	solve(nums, len(nums), &res)
	return res
}

func solve(nums []int, n int, res *[][]int) {
	if n < 2 {
		fmt.Println(nums)
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
	} else {
		for i := n - 1; i >= 0; i-- {
			tmp := nums[i]
			nums[i] = nums[n-1]
			nums[n-1] = tmp
			solve(nums, n-1, res)
			tmp = nums[i]
			nums[i] = nums[n-1]
			nums[n-1] = tmp
		}
	}
}
