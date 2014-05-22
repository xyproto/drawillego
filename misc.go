package drawille

// modulo where negative numbers gives the same result as in python
func mod(x, width int) int {
	if x < 0 {
		// make x positive while keeping the same modulo result
		x += (int(-x/width) + 1) * width
	}
	return x % width
}
