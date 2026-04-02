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
	// return a - b
}

func TestDivide(t *testing.T) {
	// return a / b
}

func TestMultiply(t *testing.T) {
	// return a * b
}
