package main

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"
)

func main() {

	f, err := os.Open("../00_images/78771293.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, err := jpeg.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("IMAGE TYPE: %T \n", img)

	bounds := img.Bounds()

	fmt.Println("Width x Height: ", bounds.Dx(), " x ", bounds.Dy())
	fmt.Println("Total Pixels: ", bounds.Dx()*bounds.Dy())

	r, g, b, a := img.At(0, 0).RGBA()
	fmt.Println("First pixel:", r, g, b, a)

	r, g, b, a = img.At(bounds.Dx(), bounds.Dy()).RGBA()
	fmt.Println("Last pixel:", r, g, b, a)
}
