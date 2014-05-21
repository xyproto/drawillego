package main

import (
	"fmt"
	. "github.com/xyproto/drawillego"
	"math"
)

const (
	rad = math.Pi / 180.0
)

func round(x float64) int {
	return int(x + 0.5)
}

func main() {
	c := NewCanvas()

	for x := 0; x < 1800; x++ {
		c.Set(x/10, round(math.Sin(float64(x)*rad*10.0)))
	}

	fmt.Println(c.Frame())

	c.Clear()

	for x := 0; x < 1800; x += 10 {
		c.Set(x/10, 10+round(math.Sin(float64(x)*rad*10.0)))
		c.Set(x/10, 10+round(math.Cos(float64(x)*rad*10.0)))
	}

	fmt.Println(c.Frame())

	c.Clear()

	for x := 0; x < 3600; x += 20 {
		c.Set(x/20, 4+round(math.Sin(float64(x)*rad*4.0)))
	}

	fmt.Println(c.Frame())

	c.Clear()

	for x := 0; x < 360; x += 4 {
		c.Set(x/4, 30+round(math.Sin(float64(x)*rad*30.0)))
	}

	for x := 0; x < 30; x++ {
		for y := 0; y < 30; y++ {
			c.Set(x, y)
			c.Toggle(x+30, y+30)
			c.Toggle(x+60, y)
		}
	}

	fmt.Println(c.Frame())
}
