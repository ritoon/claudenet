package calculate

import "testing"

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
		got := Add(d.valueInA, d.valueInB)
		if got != d.expected {
			t.Errorf("test %q: got %d, expected %d", d.testTitle, got, d.expected)
		}
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
		got := Sub(d.valueInA, d.valueInB)
		if got != d.expected {
			t.Errorf("test %q: got %d, expected %d", d.testTitle, got, d.expected)
		}
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
		got := Divide(d.valueInA, d.valueInB)
		if got != d.expected {
			t.Errorf("test %q: got %d, expected %d", d.testTitle, got, d.expected)
		}
	}
}

func TestDividePanicsOnZero(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic on divide by zero")
		}
	}()
	Divide(1, 0)
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
		got := Multiply(d.valueInA, d.valueInB)
		if got != d.expected {
			t.Errorf("test %q: got %d, expected %d", d.testTitle, got, d.expected)
		}
	}
}
