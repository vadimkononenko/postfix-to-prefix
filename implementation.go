package lab2

import (
	"errors"
	"strings"
)

func PostfixToPrefix(input string) (string, error) {
	postfixArray := strings.Split(input, " ")

	stack := make([]string, 0)

	var operatorCount int
	var operandCount int
	for _, element := range postfixArray {
		if element == "+" || element == "-" || element == "*" || element == "/" || element == "^" {
			operatorCount++
		} else {
			operandCount++
		}
	}

	if operatorCount != operandCount-1 {
		return "", errors.New("invalid prefix equation")
	}

	for i := 0; i <= len(postfixArray)-1; i++ {
		if postfixArray[i] != "+" && postfixArray[i] != "-" && postfixArray[i] != "*" && postfixArray[i] != "/" && postfixArray[i] != "^" {
			stack = append(stack, postfixArray[i])
		} else {
			if len(stack) >= 2 {
				num1 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				num2 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, postfixArray[i]+num2+num1)
			} else {
				return "", errors.New("invalid prefix equation")
			}
		}
	}

	//for i := 0; i <= len(postfixArray)-1; i++ {
	// if isOperator(postfixArray[i]) {
	//  op1 := stack[len(stack)-1]
	//  stack = stack[:len(stack)-1]
	//  op2 := stack[len(stack)-1]
	//  stack = stack[:len(stack)-1]
	//
	//  temp := postfixArray[i] + op2 + op1
	//  stack = append(stack, temp)
	// } else {
	//  stack = append(stack, postfixArray[i])
	// }
	//}

	return stack[0], nil
}
