package main

import (
	"image/jpeg"
	"log"
	"os"
	"image"
	"image/color"
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

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	nf, err := os.Create("newFile.jpg")
	if err != nil {
		log.Fatalln("my program broke creating: ", err.Error())
	}
	defer nf.Close()

	m := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			h := uint8(float64(r/255)*0.2126)
			i := uint8(float64(g/255)*0.7152)
			j := uint8(float64(b/255)*0.0722)
			k := uint8(a/255)
			m.Set(x, y, color.RGBA{h, i, j, k})
		}
	}

	o := &jpeg.Options{70}
	jpeg.Encode(nf, m, o)

}
