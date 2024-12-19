package evaluator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func EvaluateExpression(expression string) (string, error) {
	expression = strings.ReplaceAll(expression, " ", "")

	if !isValidExpression(expression) {
		return "", errors.New("invalid expression")
	}

	result, err := eval(expression)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%g", result), nil
}

func isValidExpression(expr string) bool {
	for _, char := range expr {
		if !isDigit(char) && !isOperator(char) && char != '.' {
			return false
		}
	}
	return true
}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func isOperator(char rune) bool {
	return char == '+' || char == '-' || char == '*' || char == '/'
}

func eval(expr string) (float64, error) {
	tokens := strings.FieldsFunc(expr, func(r rune) bool {
		return r == '+' || r == '*'
	})
	operators := strings.FieldsFunc(expr, func(r rune) bool {
		return isDigit(r) || r == '.'
	})

	if len(tokens) == 0 || len(operators) == 0 {
		return 0, errors.New("invalid expression")
	}

	result, err := strconv.ParseFloat(tokens[0], 64)
	if err != nil {
		return 0, err
	}

	for i, operator := range operators {
		nextValue, err := strconv.ParseFloat(tokens[i+1], 64)
		if err != nil {
			return 0, err
		}

		switch operator {
		case "+":
			result += nextValue
		case "*":
			result *= nextValue
		default:
			return 0, errors.New("unsupported operator")
		}
	}

	return result, nil
}
