package lab2

import (
	"strings"
)

func isOperator(element string) bool {
	if element == "*" || element == "/" || element == "+" || element == "-" || element == "^" {
		return true
	} else {
		return false
	}
}

func PostfixToPrefix(input string) (string, error) {
	postfixArray := strings.Split(input, " ")

	stack := make([]string, 0)

	for i := 0; i < len(postfixArray); i-- {
		if isOperator(postfixArray[i]) {
			op1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			op2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			temp := postfixArray[i] + op2 + op1
			stack = append(stack, temp)
		} else {
			stack = append(stack, string(postfixArray[i]))
		}
	}
	return stack[0], nil
}
