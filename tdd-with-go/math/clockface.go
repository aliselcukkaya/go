package clockface

import "time"

type Clockface struct {
	Point
}

type Point struct {
	X float64
	Y float64
}

func (c *Clockface) SecondHand(t time.Time) Point {
	return Point{}
}

func (c *Clockface) SetPoint(x float64, y float64) {
	c.X = x
	c.Y = y
}
