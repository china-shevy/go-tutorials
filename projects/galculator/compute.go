package main

import (
	"fmt"
	"strconv"
)

type operator func(left int64, right int64) int64

func add(left, right int64) int64 {
	return left + right
}

func sub(left, right int64) int64 {
	return left - right
}

func mul(left, right int64) int64 {
	return left * right
}

func div(left, right int64) int64 {
	return left / right
}

func compute(input string) (result string) {

	// This implementation has no lexing yet.

	operators := map[string]operator{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	operatorStack := []operator{}
	operantStack := []int64{}

	// Parsing
	for _, c := range input {
		if operator, ok := operators[string(c)]; ok {
			operatorStack = append(operatorStack, operator)
		} else if string(c) != " " {
			integer, err := strconv.ParseInt(string(c), 10, 64)
			if err != nil {
				panic(err)
			}
			operantStack = append(operantStack, integer)
		}
	}
	fmt.Println(operatorStack)
	fmt.Println(operantStack)

	// Interpreting
	var operator operator
	var left, right int64
	for len(operatorStack) > 0 {
		operator, operatorStack = operatorStack[0], operatorStack[1:]
		left, operantStack = operantStack[0], operantStack[1:]
		right, operantStack = operantStack[0], operantStack[1:]
		operantStack = append([]int64{operator(left, right)}, operantStack...)
		fmt.Println(operantStack)
	}

	return strconv.FormatInt(operantStack[0], 10)
}
