package drawille

import (
	"math"
)

type FloatPair [2]float64

func max(a, b *FloatOrInt) *FloatOrInt {
	if a.Float() >= b.Float() {
		return a
	}
	return b
}

func min(a, b *FloatOrInt) *FloatOrInt {
	if a.Float() <= b.Float() {
		return a
	}
	return b
}

// subtract
func sub(a, b *FloatOrInt) *FloatOrInt {
	return NewFloat(a.Float() - b.Float())
}

func lessEqual(a, b *FloatOrInt) bool {
	return a.Float() <= b.Float()
}

func line(x1, y1, x2, y2 int) chan FloatPair {
    return _line(NewInt(x1), NewInt(y1), NewInt(x2), NewInt(y2))
}

// Returns the float coordinates in the channel. Equivivalent to yield.
func _line(x1o, y1o, x2o, y2o *FloatOrInt) chan FloatPair {
	c := make(chan FloatPair)

	go func() {
		x1 := x1o.Normalized()
		y1 := y1o.Normalized()
		x2 := x2o.Normalized()
		y2 := y2o.Normalized()

		xdiff := sub(max(x1, x2), min(x1, x2))
		ydiff := sub(max(y1, y2), min(y1, y2))

		xdir := -1
		if lessEqual(x1, x2) {
			xdir = 1
		}

		ydir := -1
		if lessEqual(y1, y2) {
			ydir = 1
		}

		r := max(xdiff, ydiff)

		for i := 0; i < r.Int()+1; i++ {
			x := x1.Float()
			y := y1.Float()

			if ydiff.Bool() {
				y += (float64(i) * ydiff.Float()) / r.Float() * float64(ydir)
			}
			if xdiff.Bool() {
				x += (float64(i) * xdiff.Float()) / r.Float() * float64(xdir)
			}

			c <- FloatPair{x, y} // yield
		}
		close(c)

	}()

	return c
}

func Polygon(center_x, center_y float64, sides int, radius float64) chan FloatPair { // 0, 0, 4, 4
	c := make(chan FloatPair)

	go func() {

		rads := (2 * math.Pi) / float64(sides)

		for n := 0; n < sides; n++ {
			a := float64(n) * rads
			b := float64(n+1) * rads
			x1 := (center_x + math.Cos(a)) * (radius + 1) / 2.0
			y1 := (center_x + math.Sin(a)) * (radius + 1) / 2.0
			x2 := (center_x + math.Cos(b)) * (radius + 1) / 2.0
			y2 := (center_x + math.Sin(b)) * (radius + 1) / 2.0

			for fp := range _line(NewFloat(x1), NewFloat(y1), NewFloat(x2), NewFloat(y2)) {
				c <- fp // yield
			}
		}
		close(c)
	}()

	return c
}
