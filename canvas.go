package drawille

import (
	"math"
	"strings"
)

type Pos [2]int
type CharMap map[Pos]int // from position to rune, as int
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

func (c *Canvas) Set(x, y int) {
	ppos := Pos{x / 2, y / 4}
	c.chars[ppos] |= pixel_map[y%4][x%4]
}

func has(m CharMap, key Pos) bool {
	for k, _ := range m {
		if k == key {
			return true
		}
	}
	return false
}

func (c *Canvas) Unset(x, y int) {
	ppos := Pos{x / 2, y / 4}
	c.chars[ppos] &= ^pixel_map[y%4][x%4]
	if c.chars[ppos] == 0 {
		delete(c.chars, ppos)
	}
}

func (c *Canvas) Toggle(x, y int) {
	ppos := Pos{x / 2, y / 4}
	if ((c.chars[ppos] & pixel_map[y%4][x%2]) != 0) {
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
		// TODO: Might have to find the existing position instead of creating a new one
		ppos := Pos{yn, xn+i}
		c.chars[ppos] = int(b)
	}
}

func (c *Canvas) Get(x, y int) bool {
	dot_index := pixel_map[y%4][x%2]
	xn := round(float64(x) / 2.0)
	yn := round(float64(y) / 4.0)
	char, ok := c.chars[Pos{yn, xn}]
	if !ok {
		return false
	}
	return (char & dot_index) != 0
}

// characters
// min_* can be -1 for "everything"
func (c *Canvas) Rows(min_x, min_y, max_x, max_y int) (ret []int) {
	ret = []int{}

	if len(c.chars) == 0 {
		return
	}

	minrow := min_y / 4
	if min_y == -1 {
		minrow = miny(c.chars)
	}

	maxrow := (max_y - 1) / 4
	if max_y == -1 {
		maxrow = maxy(c.chars)
	}

	mincol := min_x / 2
	if min_x == -1 {
		mincol = minx(c.chars)
	}

	maxcol := (max_x -1) / 2
	if max_x == -1 {
		maxcol = maxx(c.chars)
	}

	for rownum := minrow; rownum < (maxrow + 1); rownum++ {
		if !hasy(c.chars, rownum) {
			ret = append(ret, 0)
			continue
		}
		maxcol = (max_x - 1) / 2
		if max_x == -1 {
			maxcol = maxx(c.chars)
		}
		row := []int{}

		for x := mincol; x < (maxcol + 1); x++ {
			char, found := c.chars[Pos{x, rownum}]

			if !found {
				row = append(row, 32)
			} else {
				row = append(row, braille_char_offset + char)
			}
		}

		ret = append(ret, strings.Join("", row))
	}

	return
}

func (c *Canvas) Frame(min_x, min_y, max_x, max_y int) string {
	ret = strings.Join(c.line_ending, c.Rows)
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
