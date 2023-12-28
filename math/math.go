package math

// Add returns the sum of two integers.
func Add(a, b int) int {
	return a + b
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
