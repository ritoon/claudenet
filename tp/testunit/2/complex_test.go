package complex

import (
	"testing"
)

func TestMyComplexFunction(t *testing.T) {
	data := []struct {
		testTitle string
		valueInA  int
		valueInB  int
		expected  int
	}{
		{"a", 1, 2, 1},
	}
	for _, d := range data {
		got := MyComplexFunction(d.valueInA, d.valueInB)
		if got != d.expected {
			t.Errorf("test %q: got %d, expected %d", d.testTitle, got, d.expected)
		}
	}
}
