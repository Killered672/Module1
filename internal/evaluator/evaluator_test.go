package evaluator

import (
	"testing"
)

func TestEvaluateExpression(t *testing.T) {
	tests := []struct {
		name        string
		expression  string
		expected    string
		expectedErr string
	}{
		{
			name:        "Simple addition",
			expression:  "2+2",
			expected:    "4",
			expectedErr: "",
		},
		{
			name:        "Simple multiplication",
			expression:  "3*4",
			expected:    "12",
			expectedErr: "",
		},
		{
			name:        "Complex expression",
			expression:  "2+2*2",
			expected:    "6",
			expectedErr: "",
		},
		{
			name:        "Invalid expression",
			expression:  "2+a",
			expected:    "",
			expectedErr: "invalid expression",
		},
		{
			name:        "Division by zero",
			expression:  "2/0",
			expected:    "",
			expectedErr: "division by zero",
		},
		{
			name:        "Empty expression",
			expression:  "",
			expected:    "",
			expectedErr: "invalid expression",
		},
		{
			name:        "Whitespace expression",
			expression:  " 2 + 2 ",
			expected:    "4",
			expectedErr: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := EvaluateExpression(tt.expression)

			if tt.expectedErr != "" {
				if err == nil || err.Error() != tt.expectedErr {
					t.Errorf("Expected error '%s', got '%v'", tt.expectedErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result != tt.expected {
					t.Errorf("Expected result '%s', got '%s'", tt.expected, result)
				}
			}
		})
	}
}

func TestIsValidExpression(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		expected   bool
	}{
		{
			name:       "Valid expression",
			expression: "2+2*2",
			expected:   true,
		},
		{
			name:       "Invalid characters",
			expression: "2+a",
			expected:   false,
		},
		{
			name:       "Empty expression",
			expression: "",
			expected:   true, // Пустая строка считается валидной
		},
		{
			name:       "Whitespace expression",
			expression: " 2 + 2 ",
			expected:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidExpression(tt.expression)
			if result != tt.expected {
				t.Errorf("Expected isValidExpression('%s') to be %v, got %v", tt.expression, tt.expected, result)
			}
		})
	}
}

func TestIsDigit(t *testing.T) {
	tests := []struct {
		name     string
		char     rune
		expected bool
	}{
		{
			name:     "Digit 0",
			char:     '0',
			expected: true,
		},
		{
			name:     "Digit 9",
			char:     '9',
			expected: true,
		},
		{
			name:     "Non-digit character",
			char:     'a',
			expected: false,
		},
		{
			name:     "Whitespace",
			char:     ' ',
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isDigit(tt.char)
			if result != tt.expected {
				t.Errorf("Expected isDigit('%c') to be %v, got %v", tt.char, tt.expected, result)
			}
		})
	}
}

func TestIsOperator(t *testing.T) {
	tests := []struct {
		name     string
		char     rune
		expected bool
	}{
		{
			name:     "Addition",
			char:     '+',
			expected: true,
		},
		{
			name:     "Multiplication",
			char:     '*',
			expected: true,
		},
		{
			name:     "Division",
			char:     '/',
			expected: true,
		},
		{
			name:     "Non-operator character",
			char:     'a',
			expected: false,
		},
		{
			name:     "Whitespace",
			char:     ' ',
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isOperator(tt.char)
			if result != tt.expected {
				t.Errorf("Expected isOperator('%c') to be %v, got %v", tt.char, tt.expected, result)
			}
		})
	}
}
