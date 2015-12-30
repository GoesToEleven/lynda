package main

import (
	"fmt"
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




	r, g, b, a := img.At(0, 0).RGBA()
	fmt.Println("First pixel:", r, g, b, a)

	r, g, b, a = img.At(img.Bounds().Dx(), img.Bounds().Dy()).RGBA()
	fmt.Println("Last pixel:", r, g, b, a)


	nf, err := os.Create("newFile5.jpg")
	if err != nil {
		log.Fatalln("my program broke creating: ", err.Error())
	}
	defer nf.Close()

	m := image.NewRGBA(image.Rect(0, 0, 640, 480))
	bounds := m.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			m.Set(x, y, color.RGBA{255,192,203, 255})
		}
	}

	o := &jpeg.Options{50}
	jpeg.Encode(nf, m, o)

}
