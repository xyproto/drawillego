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
		t.Errorf("No 0,0 in canvas!\n");
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
		t.Errorf("No 0,0 in canvas!\n");
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
		t.Errorf("No 0,0 in canvas!\n");
	}
	c.Toggle(0, 0)
	if len(c.chars) != 0 {
		t.Errorf("Length should be 0!\n")
	}
}

func TestSetText(t *testing.T) {
	c = NewCanvas()
	c.SetText(0, 0, "asdf")
	if c.Frame() != "asdf" {
		t.Errorf("Frame should be asdf!\n")
	}
}


/*

    def test_set_text(self):
        c = Canvas()
        c.set_text(0, 0, "asdf")
        self.assertEqual(c.frame(), "asdf")


    def test_frame(self):
        c = Canvas()
        self.assertEqual(c.frame(), '')
        c.set(0, 0)
        self.assertEqual(c.frame(), '⠁')


    def test_max_min_limits(self):
        c = Canvas()
        c.set(0, 0)
        self.assertEqual(c.frame(min_x=2), '')
        self.assertEqual(c.frame(max_x=0), '')


    def test_get(self):
        c = Canvas()
        self.assertEqual(c.get(0, 0), False)
        c.set(0, 0)
        self.assertEqual(c.get(0, 0), True)
        self.assertEqual(c.get(0, 1), False)
        self.assertEqual(c.get(1, 0), False)
        self.assertEqual(c.get(1, 1), False)


class LineTestCase(TestCase):


    def test_single_pixel(self):
        self.assertEqual(list(line(0, 0, 0, 0)), [(0, 0)])


    def test_row(self):
        self.assertEqual(list(line(0, 0, 1, 0)), [(0, 0), (1, 0)])


    def test_column(self):
        self.assertEqual(list(line(0, 0, 0, 1)), [(0, 0), (0, 1)])


    def test_diagonal(self):
        self.assertEqual(list(line(0, 0, 1, 1)), [(0, 0), (1, 1)])


	*/
