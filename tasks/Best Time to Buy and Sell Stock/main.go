package main

import "math"

func main() {
	println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	println(maxProfit([]int{7, 6, 4, 3, 1}))
}

func maxProfit(prices []int) int {
	min := math.MaxInt32
	res := 0
	for _, price := range prices {
		if min > price {
			min = price
		}

		if price-min > res {
			res = price - min
		}
	}
	return res
}
