package shapes

import "math"

type Circle struct {
	Radius float64
}

func NewCircle() Circle {
	return Circle{}
}

type Square struct {
	Length float64
	Width  float64
}

func NewSquare() Square {
	return Square{}

}

type Shape struct {
	Circle *Circle
	Square *Square
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (s *Square) Area() float64 {
	return s.Length * s.Width
}
