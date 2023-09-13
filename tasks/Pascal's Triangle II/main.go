package main

import "fmt"

func main() {
	fmt.Println(generate(5))
}

func generate(numRows int) [][]int {
	res := make([][]int, 0)
	for i := 1; i < numRows+1; i++ {
		row := make([]int, i)
		row[0] = 1
		row[i-1] = 1
		if i > 2 {
			for j := 1; j < len(res[i-2]); j++ {
				row[j] = res[i-2][j] + res[i-2][j-1]
			}
		}
		res = append(res, row)
	}
	return res
}
