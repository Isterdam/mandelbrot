package main

import (
	"image/color"
	"math/cmplx"
)

func makeCoordinates(xFrom int, colors []color.RGBA64, c chan []Coordinate) {
	var coords []Coordinate

	for x := xFrom; x < (Width/Threads)+xFrom; x++ {
		for y := 0; y < Height; y++ {
			c := complex(normalize(float64(x), true), normalize(float64(y), false)) // boolean denotes x/y
			n := divergenceRate(c)
			coords = append(coords, Coordinate{x, y, colors[n]})
		}
	}

	c <- coords
}

func divergenceRate(c complex128) int {
	fst := c
	for n := 0; n < EvalTo; n++ {
		if cmplx.Abs(c) > 2 {
			return n
		}
		c = cmplx.Pow(c, 2) + fst
	}
	return EvalTo - 1
}

// we want to normalize x and y to be in the interval [-2,1] and [-1,1] respectively
func normalize(x float64, isX bool) float64 {
	if isX {
		return ((x * 3) / Width) - 2
	}
	return ((x * 2) / Height) - 1
}
