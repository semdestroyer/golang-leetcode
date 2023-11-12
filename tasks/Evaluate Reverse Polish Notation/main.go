package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}

func evalRPN(tokens []string) int {
	numbers := make([]int, 0)
	//res := 0
	res, _ := strconv.Atoi(tokens[0])
	tokens = tokens[1:]
	for _, token := range tokens {
		fmt.Println(numbers)
		val, err := strconv.Atoi(token)
		//	fmt.Println(val, " ", err)
		if err != nil {
			switch token {
			case "+":
				term := numbers[0]
				numbers = numbers[1:]
				res += term
			case "-":
				term := numbers[0]
				numbers = numbers[1:]
				res -= term
			case "*":
				term := numbers[0]
				numbers = numbers[1:]
				res *= term
			case "/":
				term := numbers[0]
				numbers = numbers[1:]
				res /= term
			}
		} else {
			numbers = append(numbers, val)
		}
	}
	return res
}
