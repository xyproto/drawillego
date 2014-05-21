package drawille

import (
	"testing"
)

var (
	c *Canvas
)

func TestSet(t *testing.T) {
	c = NewCanvas()
	c.Set(0, 0)
	_, found := c.chars[Pos{0, 0}]
	if !found {
		t.Errorf("No 0,0 in canvas!\n")
	}
	val, found := c.chars[Pos{1, 1}]
	if found {
		t.Errorf("Found %v when there should be none!\n", val)
	}
}

func TestUnsetEmpty(t *testing.T) {
	c = NewCanvas()
	c.Set(1, 1)
	c.Unset(1, 1)
	_, found := c.chars[Pos{1, 1}]
	if found {
		t.Errorf("Found pixel where there should be none!\n")
	}
}

func TestUnsetNonempty(t *testing.T) {
	c = NewCanvas()
	c.Set(0, 0)
	c.Set(0, 1)
	c.Unset(0, 1)
	_, found := c.chars[Pos{0, 0}]
	if !found {
		t.Errorf("No 0,0 in canvas!\n")
	}
}

func TestClear(t *testing.T) {
	c = NewCanvas()
	c.Set(1, 1)
	if len(c.chars) != 1 {
		t.Errorf("Length should be 1!\n")
	}
	c.Clear()
	if len(c.chars) != 0 {
		t.Errorf("Length should be 0!\n")
	}
}

func TestToggle(t *testing.T) {
	c = NewCanvas()
	c.Toggle(0, 0)
	_, found := c.chars[Pos{0, 0}]
	if !found {
		t.Errorf("No 0,0 in canvas!\n")
	}
	c.Toggle(0, 0)
	if len(c.chars) != 0 {
		t.Errorf("Length should be 0!\n")
	}
}

func TestSetText(t *testing.T) {
	c = NewCanvas()
	c.SetText(0, 0, "asdf")
	retval := c.Frame()
	if retval != "asdf" {
		t.Errorf("Frame should be asdf, is %s!\n", retval)
	}
}

func TestFrame(t *testing.T) {
	c = NewCanvas()
	retval := c.Frame()
	if retval != "" {
		t.Errorf("Frame should be an empty string, is %s!\n", retval)
	}
	c.Set(0, 0)
	retval = c.Frame()
	if retval != "⠁" {
		t.Errorf("Frame should be an character with a single dot, is %s!\n", retval)
	}
}

func TestMaxMinLimits(t *testing.T) {
	c = NewCanvas()
	c.Set(0, 0)
	retval := c.FrameCoord(2, -1, -1, -1)
	if retval != "" {
		t.Errorf("Frame with min_x=2 should be an empty string, is %s!\n", retval)
	}
	retval = c.FrameCoord(-1, -1, 0, -1)
	if retval != "" {
		t.Errorf("Frame with max_x=0 should be an empty string, is %s!\n", retval)
	}
}

func TestGet(t *testing.T) {
	c = NewCanvas()
	if c.Get(0, 0) {
		t.Errorf("Expected get of empty position to return false.\n")
	}
	c.Set(0, 0)
	if !c.Get(0, 0) {
		t.Errorf("Expected get of 0, 0 to return true.\n")
	}
	if c.Get(0, 1) {
		t.Errorf("Expected get of 0, 1 to return false.\n")
	}
	if c.Get(1, 0) {
		t.Errorf("Expected get of 1, 0 to return false.\n")
	}
	if c.Get(1, 1) {
		t.Errorf("Expected get of 1, 1 to return false.\n")
	}
}

func TestLine(t *testing.T) {
	for fp := range line(0, 0, 0, 0) {
		if fp != (FloatPair{0.0, 0.0}) {
			t.Errorf("Expected the first coordinate on the first line to be 0.0, 0.0.\n")
		}
		break
	}

	fps := make([]FloatPair, 0, 0)
	for fp := range line(0, 0, 1, 0) {
		fps = append(fps, fp)
	}
	if fps[0] != (FloatPair{0.0, 0.0}) {
		t.Errorf("Expected the first coordinate on the second line to be 0.0, 0.0.\n")
	}
	if fps[1] != (FloatPair{1.0, 0.0}) {
		t.Errorf("Expected the second coordinate on the second line to be 1.0, 0.0.\n")
	}

	fps = make([]FloatPair, 0, 0)
	for fp := range line(0, 0, 0, 1) {
		fps = append(fps, fp)
	}
	if fps[0] != (FloatPair{0.0, 0.0}) {
		t.Errorf("Expected the first coordinate on the third line to be 0.0, 0.0.\n")
	}
	if fps[1] != (FloatPair{0.0, 1.0}) {
		t.Errorf("Expected the second coordinate on the third line to be 0.0, 1.0.\n")
	}

	fps = make([]FloatPair, 0, 0)
	for fp := range line(0, 0, 1, 1) {
		fps = append(fps, fp)
	}
	if fps[0] != (FloatPair{0.0, 0.0}) {
		t.Errorf("Expected the first coordinate on the third line to be 0.0, 0.0.\n")
	}
	if fps[1] != (FloatPair{1.0, 1.0}) {
		t.Errorf("Expected the second coordinate on the third line to be 1.0, 1.0.\n")
	}
}
