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

	r, g, b, a := img.At(0, 0).RGBA()
	fmt.Printf("%T %T %T %T \n", r, g, b, a) // uint32 uint32 uint32 uint32
	fmt.Printf("%v %v %v %v \n", r, g, b, a) // 30884 27783 21662 65535
	fmt.Printf("%d %d %d %d \n", r, g, b, a) // 30884 27783 21662 65535
	fmt.Printf("%b %b %b %b \n", r, g, b, a) // 111 1000 1010 0100 110110010000111 101010010011110 1111111111111111
	fmt.Printf("%x %x %x %x \n", r, g, b, a) // 78a4 6c87 549e ffff
}

/*
images are made up of pixels
pixels store color information:
color-depth = 1: 2 colors, monochrome
color-depth = 2: 4 colors
color-depth = 4: 16 colors
color-depth = 8: 256 colors
color-depth = 24: 16,777,216 colors, "truecolor"

2^16 = 65536

bit depth per pixel:
2 bytes each for r, g, b, a
16 + 16 + 16 + 16 = 64 bits per pixel

wikipedia:
"Images can have 64-bit pixels with 48-bit color and a 16-bit alpha channel."
https://en.wikipedia.org/wiki/Color_depth

*/
