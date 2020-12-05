package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	"time"
)

func picturize(coords *[]Coordinate, name string) {
	f, err := os.Create(name + ".png")
	if err != nil {
		fmt.Printf("Something went wrong with creating the image!")
	}

	upperLeft := image.Point{0, 0}
	upperRight := image.Point{Width, Height}
	img := image.NewRGBA(image.Rectangle{upperLeft, upperRight})

	for _, coord := range *coords {
		img.Set(coord.x, coord.y, coord.color)
	}

	png.Encode(f, img)
}

func makePalette() []color.RGBA64 {
	var colors []color.RGBA64
	rand.Seed(time.Now().UTC().Unix())

	for i := 1; i <= EvalTo; i++ {
		var r, g, b uint16
		if i > int(math.Floor(EvalTo/1.5)) { // set should be black
			r = uint16(0)
			g = uint16(0)
			b = uint16(0)
		} else {
			r = uint16(rand.Intn(65535))
			g = uint16(rand.Intn(65535))
			b = uint16(rand.Intn(65535))
		}
		colors = append(colors, color.RGBA64{r, g, b, 65535})
	}
	
	return colors
}
