package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	img := loadImage("../00_images/78771293.jpg")
	r, g, b, a := img.At(0, 0).RGBA()
	fmt.Printf("%d %d %d %d \n", r, g, b, a) // 30884 27783 21662 65535
}

func loadImage(filename string) image.Image {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, err := jpeg.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	return img
}
