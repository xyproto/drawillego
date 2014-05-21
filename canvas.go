package drawille

import (
	"math"
)

type CharMap map[int]map[int]rune
type FloatPair [2]float64

type Canvas struct {
	line_ending string
	chars       CharMap
}

func NewCanvas() *Canvas {
	return &Canvas{"\n", make(CharMap)}
}

func (c *Canvas) Clear() {
	c.chars = make(CharMap)
}

func (c *Canvas) _get_pos(xi, yi int) (*Number, *Number, int, int) {
	x := NewInt(xi)
	y := NewInt(yi)
	cols := x.Int() / 2
	rows := y.Int() / 4
	return x, y, cols, rows
}

func intOr(a, b *Number) *Number {
	return NewInt(a.Int() | b.Int())
}

func intAndUnaryBitwiseComplement(a, b *Number) *Number {
	return NewInt(a.Int() & ^b.Int())
}

func (c *Canvas) Set(xi, yi int) {
	_, _, px, py := c._get_pos(xi, yi)
	val := c.chars[py][px]
	newval := intOr(NewInt(int(val)), pixel_map[yi%4][xi%2])
	c.chars[py][px] = rune(newval.Int())
}

func has(m CharMap, key int) bool {
	for k, _ := range m {
		if k == key {
			return true
		}
	}
	return false
}

func (c *Canvas) Unset(xi, yi int) {
	x, y, px, py := c._get_pos(xi, yi)
	val := c.chars[py][px]
	if val.floatType == false {
		newval := intAndUnaryBitwiseComplement(&val, pixel_map[y.Int()%4][x.Int()%2])
		c.chars[py][px] = *newval
	}
	val = c.chars[py][px]
	if val.floatType || (val.Int() == 0) {
		delete(c.chars[py], px)
	}
	if has(c.chars, py) {
		delete(c.chars, py)
	}
}

func (c *Canvas) Toggle(x, y int) {
	_, _, px, py := c._get_pos(x, y)
	val := c.chars[py][px]
	if val.floatType || (intOr(&val, pixel_map[y%4][x%2]).Int() != 0) {
		c.Unset(x, y)
	} else {
		c.Set(x, y)
	}
}

func round(a float64) int {
	return int(a + 0.5)
}

func (c *Canvas) SetText(x, y int, text string) {
	xn := round(float64(x) / 2.0)
	yn := round(float64(y) / 4.0)
	for i, b := range text {
		newval := NewInt(int(b))
		c.chars[yn][xn+i] = *newval
	}
}

func (c *Canvas) Get(x, y *Number) bool {
	dot_index := pixel_map[y.Int()%4][x.Int()%2]
	xn := normalize(NewFloat(x.Float() / 2.0))
	yn := normalize(NewFloat(y.Float() / 4.0))
	char, ok := c.chars[yn.Int()][xn.Int()]
	if !ok {
		return false
	}
	if char.floatType {
		return true
	}
	return (char.Int() & dot_index.Int()) != 0
}

func (c *Canvas) Rows(min_x, min_y, max_x, max_y *Number) []*Number {
	return []*Number{}
}

func (c *Canvas) Frame(min_x, min_y, max_x, max_y *Number) string {
	return ""
}

func max(a, b *Number) *Number {
	if a.Float() >= b.Float() {
		return a
	}
	return b
}

func min(a, b *Number) *Number {
	if a.Float() <= b.Float() {
		return a
	}
	return b
}

// subtract
func sub(a, b *Number) *Number {
	return NewFloat(a.Float() - b.Float())
}

func lessEqual(a, b *Number) bool {
	return a.Float() <= b.Float()
}


// Returns the float coordinates in the channel. Equivivalent to yield.
func line(x1o, y1o, x2o, y2o *Number) chan FloatPair {
	c := make(chan FloatPair)

	go func() {

		x1 := normalize(x1o)
		y1 := normalize(y1o)
		x2 := normalize(x2o)
		y2 := normalize(y2o)

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

			for fp := range line(NewFloat(x1), NewFloat(y1), NewFloat(x2), NewFloat(y2)) {
				c <- fp // yield
			}
		}
		close(c)
	}()

	return c
}
