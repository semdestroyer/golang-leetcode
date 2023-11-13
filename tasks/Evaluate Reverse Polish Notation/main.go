package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(13 / 5)
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for i := range tokens {
		val, err := strconv.Atoi(tokens[i])
		if err != nil {
			term1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			term2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			switch tokens[i] {
			case "+":
				stack = append(stack, term1+term2)
			case "-":
				stack = append(stack, term1-term2)
			case "*":
				stack = append(stack, term1*term2)
			case "/":
				stack = append(stack, term2/term1)
			}

		} else {
			stack = append(stack, val)
		}
		fmt.Println(stack)
	}
	res := stack[0]
	return res
}
