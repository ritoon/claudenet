package complex

func MyComplexFunction(a, b int) int {
	if a < b {
		return b - 1
	}
	// la flèche de la mort: death arrow
	if a != b {
		if a == 0 {
			if b == -1 {
				return b + 2
			}
			return (a + 10) * b
		}
		return a / 2
	}
	return b % 2
}
