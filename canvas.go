package drawille

type (
	Pos            [2]int
	CharMap        map[Pos]rune
	RegularTextMap map[Pos]bool
	Row            []rune

	Canvas struct {
		line_ending string
		chars       CharMap
		regular     RegularTextMap
	}
)

func NewCanvas() *Canvas {
	return &Canvas{"\n", make(CharMap), make(RegularTextMap)}
}

func (c *Canvas) Clear() {
	c.chars = make(CharMap)
}

func (c *Canvas) Set(x, y int) {
	ppos := Pos{x / 2, y / 4}
	// Skip regular text
	if regular, ok := c.regular[ppos]; ok && regular {
		return
	}
	// Set the correct dot pattern
	c.chars[ppos] |= int32(pixel_map[y%4][x%2])
}

func hasx(m CharMap, xkey int) bool {
	for k, _ := range m {
		if k[0] == xkey {
			return true
		}
	}
	return false
}

func hasy(m CharMap, ykey int) bool {
	for k, _ := range m {
		if k[1] == ykey {
			return true
		}
	}
	return false
}

func minx(m CharMap) int {
	mx := -1
	if len(m) > 0 {
		// Fetch the first x value
		for k, _ := range m {
			mx = k[0]
			break
		}
	} else {
		return mx
	}
	// Find the smallest x value
	for k, _ := range m {
		if k[0] < mx {
			mx = k[0]
		}
	}
	return mx
}

func miny(m CharMap) int {
	my := -1
	if len(m) > 0 {
		// Fetch the first y value
		for k, _ := range m {
			my = k[1]
			break
		}
	} else {
		return my
	}
	// Find the smallest y value
	for k, _ := range m {
		if k[1] < my {
			my = k[1]
		}
	}
	return my
}

func maxx(m CharMap) int {
	mx := -1
	if len(m) > 0 {
		// Fetch the first x value
		for k, _ := range m {
			mx = k[0]
			break
		}
	} else {
		return mx
	}
	// Find the largest x value
	for k, _ := range m {
		if k[0] > mx {
			mx = k[0]
		}
	}
	return mx
}

func maxy(m CharMap) int {
	my := -1
	if len(m) > 0 {
		// Fetch the first y value
		for k, _ := range m {
			my = k[1]
			break
		}
	} else {
		return my
	}
	// Find the largest y value
	for k, _ := range m {
		if k[1] > my {
			my = k[1]
		}
	}
	return my
}

func (c *Canvas) Unset(x, y int) {
	ppos := Pos{x / 2, y / 4}
	c.chars[ppos] &= ^pixel_map[y%4][x%2]
	if c.chars[ppos] == 0 {
		delete(c.chars, ppos)
	}
	if _, ok := c.regular[ppos]; ok {
		delete(c.regular, ppos)
	}
}

func (c *Canvas) Toggle(x, y int) {
	ppos := Pos{x / 2, y / 4}
	if (c.chars[ppos] & pixel_map[y%4][x%2]) != 0 {
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
	var ppos Pos
	for i, ru := range text {
		ppos = Pos{yn, xn + i}
		// set the rune
		c.chars[ppos] = ru
		// mark as regular text
		c.regular[ppos] = true
	}
}

func (c *Canvas) Get(x, y int) bool {
	dot_index := pixel_map[y%4][x%2]
	xn := round(float64(x) / 2.0)
	yn := round(float64(y) / 4.0)
	ppos := Pos{yn, xn}
	char, ok := c.chars[ppos]
	if !ok {
		return false
	}
	// Regular text
	if regular, ok := c.regular[ppos]; ok && regular {
		return true
	}
	return (char & dot_index) != 0
}

// characters
// min_* can be -1 for "everything"
func (c *Canvas) Rows(min_x, min_y, max_x, max_y int) (ret []Row) {
	if len(c.chars) == 0 {
		return
	}

	//log.Println("min_x", min_x, "min_y", min_y, "max_x", max_x, "max_y", max_y)

	var (
		minrow, maxrow, mincol, maxcol int
		ppos                           Pos
	)

	if min_y != -1 {
		minrow = min_y / 4
	} else {
		minrow = miny(c.chars)
	}

	if max_y != -1 {
		maxrow = (max_y - 1) / 4
	} else {
		maxrow = maxy(c.chars)
	}

	if min_x != -1 {
		mincol = min_x / 2
	} else {
		mincol = minx(c.chars)
	}

	if max_x != -1 {
		maxcol = (max_x - 1) / 2
	} else {
		maxcol = maxx(c.chars)
	}

	for y := minrow; y < (maxrow + 1); y++ {

		//log.Printf("y %d, from %d to %d\n", y, minrow, (maxrow + 1))

		if !hasy(c.chars, y) {
			ret = append(ret, Row{})
			continue
		}
		if max_x != -1 {
			maxcol = (max_x - 1) / 2
			//log.Println("given maxcol", maxcol)
		} else {
			maxcol = maxx(c.chars)
			//log.Println("found maxcol", maxcol)
		}
		row := []rune{}

		for x := mincol; x < (maxcol + 1); x++ {

			//log.Printf("x %d, from %d to %d\n", x, mincol, (maxcol + 1))

			ppos = Pos{x, y}
			char, found := c.chars[ppos]

			if !found {
				row = append(row, 32)
			} else if regular, ok := c.regular[ppos]; ok && regular {
				row = append(row, char)
			} else {
				row = append(row, int32(braille_char_offset)+int32(char))
			}
		}

		ret = append(ret, row)
	}

	return
}

func (c *Canvas) Frame() string {
	return c.FrameCoord(-1, -1, -1, -1)
}

func (c *Canvas) FrameCoord(min_x, min_y, max_x, max_y int) string {
	s := ""
	for _, row := range c.Rows(min_x, min_y, max_x, max_y) {
		for _, rowChar := range row {
			s += string(rowChar)
		}
	}
	return s
}
