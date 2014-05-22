package main

import (
	. "github.com/nsf/termbox-go"
	. "github.com/xyproto/drawillego"
	"github.com/xyproto/textgui"
	"math"
	"os"
	"strings"
	"time"
)

const RAD = math.Pi / 180.0

type (
	Point3D struct {
		x float64
		y float64
		z float64
	}
	Face []int
)

var (
	vertices []Point3D = []Point3D{
		Point3D{-20.0, 20.0, -20.0},
		Point3D{20.0, 20.0, -20.0},
		Point3D{20.0, -20.0, -20.0},
		Point3D{-20.0, -20.0, -20.0},
		Point3D{-20.0, 20.0, 20.0},
		Point3D{20.0, 20.0, 20.0},
		Point3D{20.0, -20.0, 20.0},
		Point3D{-20.0, -20.0, 20.0},
	}
	faces []Face = []Face{
		Face{0, 1, 2, 3},
		Face{1, 5, 6, 2},
		Face{5, 4, 7, 6},
		Face{4, 0, 3, 7},
		Face{0, 4, 5, 1},
		Face{3, 2, 6, 7},
	}
)

func NewPoint3D(x, y, z float64) *Point3D {
	return &Point3D{x, y, z}
}

func (p *Point3D) RotateX(angle float64) *Point3D {
	rad := RAD * angle
	cosa := math.Cos(rad)
	sina := math.Sin(rad)
	y := p.y*cosa - p.z*sina
	z := p.y*sina + p.z*cosa
	return &Point3D{p.x, y, z}
}

func (p *Point3D) RotateY(angle float64) *Point3D {
	rad := RAD * angle
	cosa := math.Cos(rad)
	sina := math.Sin(rad)
	z := p.z*cosa - p.x*sina
	x := p.z*sina + p.x*cosa
	return &Point3D{x, p.y, z}
}

func (p *Point3D) RotateZ(angle float64) *Point3D {
	rad := RAD * angle
	cosa := math.Cos(rad)
	sina := math.Sin(rad)
	x := p.x*cosa - p.y*sina
	y := p.x*sina + p.y*cosa
	return &Point3D{x, y, p.z}
}

func (p *Point3D) Project(win_width, win_height, fov, viewer_distance float64) *Point3D {
	factor := fov / (viewer_distance + p.z)
	x := p.x*factor + win_width/2.0
	y := -p.y*factor + win_height/2.0
	return &Point3D{x, y, 1.0}
}

func run(projection bool) {
	var t []Point3D
	var p *Point3D

	angleX, angleY, angleZ := 0.0, 0.0, 0.0
	c := NewCanvas()
	for {
		//for rounds := 0; rounds < 1000; rounds++ {

		// Will hold transformed vertices.
		t = []Point3D{}

		for _, v := range vertices {
			// Rotate the point around X axis, then around Y axis, and finally around Z axis.
			p = &v
			p = p.RotateX(angleX)
			p = p.RotateY(angleY)
			p = p.RotateZ(angleZ)
			if projection {
				// Transform the point from 3D to 2D
				p = p.Project(50, 50, 50, 50)
			}
			// Put the point in the list of transformed vertices
			t = append(t, *p)
		}

		for _, f := range faces {
			for fp := range Linef(t[f[0]].x, t[f[0]].y, t[f[1]].x, t[f[1]].y) {
				c.Set(int(fp[0]), int(fp[1]))
			}
			for fp := range Linef(t[f[1]].x, t[f[1]].y, t[f[2]].x, t[f[2]].y) {
				c.Set(int(fp[0]), int(fp[1]))
			}
			for fp := range Linef(t[f[2]].x, t[f[2]].y, t[f[3]].x, t[f[3]].y) {
				c.Set(int(fp[0]), int(fp[1]))
			}
			for fp := range Linef(t[f[3]].x, t[f[3]].y, t[f[0]].x, t[f[0]].y) {
				c.Set(int(fp[0]), int(fp[1]))
			}
		}

		f := c.FrameCoord(-40, -40, 80, 80)

		//stdscr.AddStr(0, 0, '{0}\n'.format(f))
		xoffset := 2
		for y, line := range strings.Split(f, "\n") {
			textgui.Write(xoffset, y, line, ColorRed|AttrBold, ColorBlack|AttrBold)
		}

		//stdscr.Refresh()
		textgui.Flush()

		angleX += 2.0
		angleY += 3.0
		angleZ += 5.0

		time.Sleep(50 * time.Millisecond)

		c.Clear()

		textgui.Clear()

		// TODO: Fork termbox and implement PeekEvent
		//e := PeekEvent()
		//switch e.Type {
		//case EventKey:
		//	switch e.Key {
		//	case KeyEsc, KeyEnter, KeySpace:
		//		return
		//	}
		//}
	}
}

func main() {
	projection := false
	if len(os.Args) > 0 {
		if os.Args[0] == "-p" {
			projection = true
		}
	}
	textgui.Init()
	run(projection)
	textgui.Close()
}
