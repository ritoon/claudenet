package complex

import "testing"

func TestMyComplexFunction(t *testing.T) {
	data := []struct {
		testTitle string
		valueInA  int
		valueInB  int
		expected  int
	}{
		{"a", 1, 2, 1},
		{"b", 2, 2, 0},
		{"c", 0, -1, 1},
		{"d", -1, -2, 0},
		{"e", 0, -3, -30},
	}
	for _, d := range data {
		got := MyComplexFunction(d.valueInA, d.valueInB)
		if got != d.expected {
			t.Errorf("test %q: got %d, expected %d", d.testTitle, got, d.expected)
		}
	}
}

func FuzzMyComplexFunction(f *testing.F) {
	corpus := []int{-10, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 10}
	for _, a := range corpus {
		for _, b := range corpus {
			f.Add(a, b)
		}
	}
	f.Fuzz(func(t *testing.T, a, b int) {
		got := MyComplexFunction(a, b)

		switch {
		case a < b:
			want := b - 1
			if got != want {
				t.Fatalf("a<b: got %d, want %d (a=%d,b=%d)", got, want, a, b)
			}
		case a == b:
			want := b % 2
			if got != want {
				t.Fatalf("a==b: got %d, want %d (a=%d,b=%d)", got, want, a, b)
			}
		default:
			if a == 0 {
				if b == -1 {
					if got != 1 {
						t.Fatalf("a!=b,a==0,b==-1: got %d, want 1 (a=%d,b=%d)", got, a, b)
					}
				} else {
					want := 10 * b
					if got != want {
						t.Fatalf("a!=b,a==0: got %d, want %d (a=%d,b=%d)", got, want, a, b)
					}
				}
			} else {
				want := a / 2
				if got != want {
					t.Fatalf("a!=b,a!=0: got %d, want %d (a=%d,b=%d)", got, want, a, b)
				}
			}
		}
	})
}
