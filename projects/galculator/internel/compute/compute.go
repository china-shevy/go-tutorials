package compute

import (
	"fmt"
	"galculator/internel/lexer"
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

func Compute(input string) (result string) {

	// This implementation has no lexing yet.

	operators := map[lexer.Operator]operator{
		lexer.Add: add,
		lexer.Sub: sub,
		lexer.Mul: mul,
		lexer.Div: div,
	}

	operatorStack := []operator{}
	operantStack := []int64{}

	// Parsing
	tokens, err := lexer.Lex(input)
	if err != nil {
		return err.Error()
	}
	for _, token := range tokens {

		switch t := token.(type) {
		case lexer.Operator:
			if operator, ok := operators[t]; ok {
				operatorStack = append(operatorStack, operator)
			} else {
				panic("?")
			}
		case lexer.Number:
			integer, err := strconv.ParseInt(t.Value, 10, 64)
			if err != nil {
				panic(err)
			}
			operantStack = append(operantStack, integer)
		}
	}

	// Interpreting
	var operator operator
	var left, right int64
	for len(operatorStack) > 0 {
		fmt.Println(operantStack)
		operator, operatorStack = operatorStack[0], operatorStack[1:]
		left, operantStack = operantStack[0], operantStack[1:]
		right, operantStack = operantStack[0], operantStack[1:]
		operantStack = append([]int64{operator(left, right)}, operantStack...)
	}

	return strconv.FormatInt(operantStack[0], 10)
}
