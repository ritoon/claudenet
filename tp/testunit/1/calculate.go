package calculate

import (
	"errors"
	"math"
)

var ErrOverflow = errors.New("integer overflow")
var ErrDivisionByZero = errors.New("division by zero")

func Add(a, b int) (int, error) {
	if (b > 0 && a > math.MaxInt-b) || (b < 0 && a < math.MinInt-b) {
		return 0, ErrOverflow
	}
	return a + b, nil
}

func Sub(a, b int) (int, error) {
	if (b < 0 && a > math.MaxInt+b) || (b > 0 && a < math.MinInt+b) {
		return 0, ErrOverflow
	}
	return a - b, nil
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	if a == math.MinInt && b == -1 {
		return 0, ErrOverflow
	}
	return a / b, nil
}

func Multiply(a, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, nil
	}
	result := a * b
	if result/a != b {
		return 0, ErrOverflow
	}
	return result, nil
}
