package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{4, 2, 3}))
	fmt.Println(trap([]int{5, 4, 1, 2}))
	fmt.Println(trap([]int{4, 9, 4, 5, 3, 2}))
	fmt.Println(trap([]int{9, 6, 8, 8, 5, 6, 3}))
}

func trap(height []int) int {
	res := 0
	lw := 0
	rw := 1
	waterArr := make([]int, len(height))

	for lw < len(height)-1 {
		if height[lw] > height[rw] {

			if height[rw] > height[rw-1] {
				for i := lw + 1; i < rw; i++ {
					waterArr[i] = int(math.Min(float64(height[rw]), float64(height[lw]))) - height[i]
					if waterArr[i] < 0 {
						waterArr[i] = 0
					}
				}
			}
			if rw >= len(height)-1 {
				lw++
				rw = lw
			} else {
				rw++
			}

		} else {
			for i := lw + 1; i < rw; i++ {
				waterArr[i] = int(math.Min(float64(height[rw]), float64(height[lw]))) - height[i]
				if waterArr[i] < 0 {
					waterArr[i] = 0
				}
			}
			lw = rw
			rw++
		}

		fmt.Println(lw, ":", rw, " h", height[lw], ":", height[rw])
	}
	fmt.Println(waterArr)
	for i := 0; i < len(waterArr); i++ {
		res += waterArr[i]
	}
	return res
}

//func trap(height []int) int {
//	res := 0
//	lw := 0
//	rw := 1
//	waterArr := make([]int, len(height))
//
//	for lw < len(height)-1 {
//
//		if height[lw] > height[rw] {
//
//			if rw == len(height)-1 {
//				if height[rw] > height[rw-1] {
//					for i := lw + 1; i < rw; i++ {
//						waterArr[i] = int(math.Min(float64(height[rw]), float64(height[lw]))) - height[i]
//						if waterArr[i] < 0 {
//							waterArr[i] = 0
//						}
//					}
//					break
//				} else {
//					lw++
//					rw = lw
//				}
//
//			} else {
//				rw++
//			}
//			//fmt.Println(lw, ":", rw, " h", height[lw], ":", height[rw])
//		} else {
//			for i := lw + 1; i < rw; i++ {
//				waterArr[i] = int(math.Min(float64(height[rw]), float64(height[lw]))) - height[i]
//			}
//			lw = rw
//			rw++
//		}
//	}
//	//fmt.Println(waterArr)
//	for i := 0; i < len(waterArr); i++ {
//		res += waterArr[i]
//	}
//	return res
//}
//
