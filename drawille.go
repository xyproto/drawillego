package drawille

// Direct port of the drawille library for python, by Adam Tauber <asciimoo@gmail.com> (GPL3)
// Port is done by Alexander Rødseth <rodseth@gmail.com>

import (
	. "github.com/xyproto/textgui"
)

/*
 * http://www.alanwood.net/unicode/braille_patterns.html
 *
 * dots:
 *   ,___,
 *   |1 4|
 *   |2 5|
 *   |3 6|
 *   |7 8|
 *   `````
 */

// For a direct translation from untyped Python to typed Go
type Numbers []*Number

var (
	pixel_map = []Numbers{Numbers{NewInt(0x01), NewInt(0x08)},
		Numbers{NewInt(0x02), NewInt(0x10)},
		Numbers{NewInt(0x04), NewInt(0x20)},
		Numbers{NewInt(0x40), NewInt(0x80)}}
)

const (
	// braille unicode characters starts at 0x2800
	braille_char_offset = 0x2800
)

// Returns terminal width, height
func getTerminalSize() (int, int) {
	return ScreenWidth(), ScreenHeight()
}

func normalize(fi *Number) *Number {
	return fi.Normalized()
}
