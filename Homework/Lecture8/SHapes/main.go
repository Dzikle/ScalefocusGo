package main

import (
	"fmt"
	"shapes/shapes"
)

type Shapes []Shape

type Shape interface {
	Area() float64
}

func (s Shapes) LargestArea() float64 {
	var result float64
	for _, s := range s {
		shape := s.Area()
		if shape > result {
			result = shape
		}
	}
	return result
}

func main() {

	circle := shapes.NewCircle()
	circle.Radius = 3.00
	square := shapes.NewSquare()
	square.Length = 7.00
	square.Width = 7.00

	var sliceSh Shapes

	sliceSh = append(sliceSh, &circle)
	sliceSh = append(sliceSh, &square)
	fmt.Println(sliceSh.LargestArea())

}
