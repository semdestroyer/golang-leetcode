package main

func main() {
	println(isValid(""))
}

func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, char := range s {
		stack = append(stack, char)
		if len(stack) < 2 {
			continue
		}
		slice := stack[len(stack)-2:]
		if string(slice) == "()" || string(slice) == "[]" || string(slice) == "{}" {
			stack = stack[:len(stack)-2]
		}
	}
	println(string(stack))
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
