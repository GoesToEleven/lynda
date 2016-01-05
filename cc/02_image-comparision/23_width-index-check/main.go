package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
)

type pixel struct {
	r, g, b, a uint32
}

func main() {

	f, err := os.Open("../00d_width-index-check/ten-by-two.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, err := jpeg.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	min := bounds.Min
	max := bounds.Max
	fmt.Println("Width:", width)
	fmt.Println("Height:", height)
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)

	r, g, b, a := img.At(0, 0).RGBA()
	fmt.Println("First pixel:", r, g, b, a)

	r, g, b, a = img.At(bounds.Dx(), bounds.Dy()).RGBA()
	fmt.Println("Last pixel:", r, g, b, a)

	xPixels := getPixels(img)
	for i, v := range xPixels {
		fmt.Println(i, v)
	}
}

func getPixels(img image.Image) []pixel {

	bounds := img.Bounds()
	pixels := make([]pixel, bounds.Dx()*bounds.Dy()+1)

	for i := 0; i <= bounds.Dx()*bounds.Dy(); i++ {
		x := i % bounds.Dx()
		y := i / bounds.Dx()
		r, g, b, a := img.At(x, y).RGBA()
		pixels[i].r = r
		pixels[i].g = g
		pixels[i].b = b
		pixels[i].a = a
	}

	return pixels
}
