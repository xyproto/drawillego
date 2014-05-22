package drawille

// Round off a positive floating point number
func Round(a float64) int {
	return int(a + 0.5)
}

// modulo where negative numbers gives the same result as in python
func Mod(x, width int) int {
	if x < 0 {
		// make x positive while keeping the same modulo result
		x += (int(-x/width) + 1) * width
	}
	return x % width
}
