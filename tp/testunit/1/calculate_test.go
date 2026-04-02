package calculate

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	data := []struct {
		testTitle string
		valueInA  int
		valueInB  int
		expected  int
	}{
		{"a", 1, 1, 2},
		{"b", 1, 9, 10},
	}
	for _, d := range data {
		got, err := Add(d.valueInA, d.valueInB)
		if err != nil {
			t.Errorf("test %q: unexpected error: %v", d.testTitle, err)
			continue
		}
		if got != d.expected {
			t.Errorf("test %q: got %d, expected %d", d.testTitle, got, d.expected)
		}
	}
}

func TestAddOverflow(t *testing.T) {
	_, err := Add(math.MaxInt, 1)
	if err != ErrOverflow {
		t.Errorf("expected ErrOverflow, got %v", err)
	}

	_, err = Add(math.MinInt, -1)
	if err != ErrOverflow {
		t.Errorf("expected ErrOverflow for MinInt + (-1), got %v", err)
	}
}

func TestSub(t *testing.T) {
	data := []struct {
		testTitle string
		valueInA  int
		valueInB  int
		expected  int
	}{
		{"positive", 10, 3, 7},
		{"negative result", 3, 10, -7},
		{"zeros", 0, 0, 0},
		{"sub zero", 5, 0, 5},
		{"negative numbers", -3, -2, -1},
	}
	for _, d := range data {
		got, err := Sub(d.valueInA, d.valueInB)
		if err != nil {
			t.Errorf("test %q: unexpected error: %v", d.testTitle, err)
			continue
		}
		if got != d.expected {
			t.Errorf("test %q: got %d, expected %d", d.testTitle, got, d.expected)
		}
	}
}

func TestSubOverflow(t *testing.T) {
	_, err := Sub(math.MinInt, 1)
	if err != ErrOverflow {
		t.Errorf("expected ErrOverflow, got %v", err)
	}

	_, err = Sub(math.MaxInt, -1)
	if err != ErrOverflow {
		t.Errorf("expected ErrOverflow for MaxInt - (-1), got %v", err)
	}
}

func TestDivide(t *testing.T) {
	data := []struct {
		testTitle string
		valueInA  int
		valueInB  int
		expected  int
	}{
		{"simple", 10, 2, 5},
		{"integer division", 7, 2, 3},
		{"divide by 1", 5, 1, 5},
		{"zero numerator", 0, 5, 0},
		{"negative", -10, 2, -5},
		{"both negative", -10, -2, 5},
	}
	for _, d := range data {
		got, err := Divide(d.valueInA, d.valueInB)
		if err != nil {
			t.Errorf("test %q: unexpected error: %v", d.testTitle, err)
			continue
		}
		if got != d.expected {
			t.Errorf("test %q: got %d, expected %d", d.testTitle, got, d.expected)
		}
	}
}

func TestDivideDivisionByZero(t *testing.T) {
	_, err := Divide(1, 0)
	if err != ErrDivisionByZero {
		t.Errorf("expected ErrDivisionByZero, got %v", err)
	}
}

func TestDivideOverflow(t *testing.T) {
	_, err := Divide(math.MinInt, -1)
	if err != ErrOverflow {
		t.Errorf("expected ErrOverflow, got %v", err)
	}
}

func TestMultiply(t *testing.T) {
	data := []struct {
		testTitle string
		valueInA  int
		valueInB  int
		expected  int
	}{
		{"positive", 3, 4, 12},
		{"by zero", 5, 0, 0},
		{"by one", 7, 1, 7},
		{"negatives", -3, -4, 12},
		{"mixed sign", -3, 4, -12},
		{"zeros", 0, 0, 0},
	}
	for _, d := range data {
		got, err := Multiply(d.valueInA, d.valueInB)
		if err != nil {
			t.Errorf("test %q: unexpected error: %v", d.testTitle, err)
			continue
		}
		if got != d.expected {
			t.Errorf("test %q: got %d, expected %d", d.testTitle, got, d.expected)
		}
	}
}

func TestMultiplyOverflow(t *testing.T) {
	_, err := Multiply(math.MaxInt, 2)
	if err != ErrOverflow {
		t.Errorf("expected ErrOverflow, got %v", err)
	}

	_, err = Multiply(math.MinInt, 2)
	if err != ErrOverflow {
		t.Errorf("expected ErrOverflow for MinInt * 2, got %v", err)
	}
}
