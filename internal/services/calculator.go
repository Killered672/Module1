package services

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Calculate(expression string) (string, error) {
	valid, err := regexp.MatchString(`^[\d\s\+\-\*/\(\)]+$`, expression)
	if err != nil || !valid {
		return "", errors.New("invalid expression")
	}

	expression = strings.ReplaceAll(expression, " ", "")

	result, err := eval(expression)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%.2f", result), nil
}

func eval(expression string) (float64, error) {
	result, err := strconv.ParseFloat(expression, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}
