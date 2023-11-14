package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}

func summaryRanges(nums []int) []string {
	ranges := make([]string, 0)
	str := ""
	for i := 0; i < len(nums); i++ {
		if i+1 < len(nums) && nums[i] == nums[i+1]-1 {
			if str == "" {
				str = strconv.Itoa(nums[i]) + "->"
			}
		} else {
			if str == "" {
				ranges = append(ranges, strconv.Itoa(nums[i]))
			} else {
				ranges = append(ranges, str+strconv.Itoa(nums[i]))
				str = ""
			}
		}
	}

	//lp := 0
	//rp := 0
	//for lp < len(nums)-2 || rp < len(nums)-2 {
	//	if nums[rp] == nums[rp+1]-1 {
	//		rp++
	//	} else {
	//		if rp == lp {
	//			ranges = append(ranges, strconv.Itoa(nums[lp])
	//		} else {
	//			ranges = append(ranges, strconv.Itoa(nums[lp])+"->"+strconv.Itoa(nums[rp]))
	//		}
	//		rp++
	//		lp = rp
	//	}
	//}

	//head := 0
	//end := 0
	//prev := 0
	//for num := range nums {
	//	if prev+1 == num {
	//		prev = num
	//		head = num
	//	} else if head != 0 {
	//		end = num
	//		ranges = append(ranges, strconv.Itoa(head)+"->"+strconv.Itoa(end))
	//	} else {
	//		ranges = append(ranges, strconv.Itoa(num))
	//	}
	//}
	return ranges
}
