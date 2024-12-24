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

	result, err := evalWithPrecedence(expression)
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

func evalWithPrecedence(expr string) (float64, error) {
	for {
		index := strings.IndexAny(expr, "*/")
		if index == -1 {
			break
		}

		left := expr[:index]
		right := expr[index+1:]

		leftTokens := strings.FieldsFunc(left, func(r rune) bool {
			return isOperator(r)
		})

		leftValue, err := strconv.ParseFloat(leftTokens[len(leftTokens)-1], 64)
		if err != nil {
			return 0, err
		}

		rightTokens := strings.FieldsFunc(right, func(r rune) bool {
			return isOperator(r)
		})

		rightValue, err := strconv.ParseFloat(rightTokens[0], 64)
		if err != nil {
			return 0, err
		}

		var result float64
		switch expr[index] {
		case '*':
			result = leftValue * rightValue
		case '/':
			if rightValue == 0 {
				return 0, errors.New("division by zero")
			}
			result = leftValue / rightValue
		}

		expr = left[:len(left)-len(leftTokens[len(leftTokens)-1])] + fmt.Sprintf("%g", result) + right[len(rightTokens[0]):]
	}

	for {
		index := strings.IndexAny(expr, "+-")
		if index == -1 {
			break
		}

		left := expr[:index]
		right := expr[index+1:]

		leftTokens := strings.FieldsFunc(left, func(r rune) bool {
			return isOperator(r)
		})

		leftValue, err := strconv.ParseFloat(leftTokens[len(leftTokens)-1], 64)
		if err != nil {
			return 0, err
		}

		rightTokens := strings.FieldsFunc(right, func(r rune) bool {
			return isOperator(r)
		})

		rightValue, err := strconv.ParseFloat(rightTokens[0], 64)
		if err != nil {
			return 0, err
		}

		var result float64
		switch expr[index] {
		case '+':
			result = leftValue + rightValue
		case '-':
			result = leftValue - rightValue
		}

		expr = left[:len(left)-len(leftTokens[len(leftTokens)-1])] + fmt.Sprintf("%g", result) + right[len(rightTokens[0]):]
	}

	result, err := strconv.ParseFloat(expr, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
