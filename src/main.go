package main

import (
	"fmt"
	"image/color"
	"os"
	"strconv"
	"time"
)

type Coordinate struct {
	x     int
	y     int
	color color.RGBA64
}

const (
	Width   = 6000
	Height  = Width
	Threads = 20 // must divide Width evenly
	EvalTo  = 64 // times that complex number is squared
)

func main() {
	var images int
	fmt.Println("Enter number of images to be generated:")
	fmt.Scan(&images)

	for i := 0; i < images; i++ {
		generate(i)
	}

	fmt.Println("\nAll images were generated successfully! Enter '0' to exit.")
	fmt.Scan(&images) // wait for user to exit terminal
}

func generate(i int) {
	colors := makePalette()
	var coords []Coordinate
	parts := make(chan []Coordinate, Threads)

	start := time.Now()

	for x := 0; x < Width; x += (Width / Threads) {
		go func(xFrom int, colors *[]color.RGBA64, c chan []Coordinate) {
			makeCoordinates(xFrom, *colors, c)
		}(x, &colors, parts)
	}

	for i := 0; i < Threads; i++ {
		part := <-parts
		coords = append(coords, part...)
	}

	picturize(&coords, "image"+strconv.Itoa(i+1))

	path, err := os.Getwd()
	if err != nil {
		fmt.Printf("Could not get working directory!")
		return
	}

	fmt.Print("Generated image number " + strconv.Itoa(i+1) + " to " + path + " in ")
	fmt.Println(time.Now().Sub(start))
}
