package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"time"
)

// at terminal:
// go run -race main.go

type pixel struct {
	r, g, b, a uint32
}

func main() {
	start := time.Now()
	chP := getPaths()
	chI := getPixels(chP)

	var i int
	for xPix := range chI {
		for j, pixel := range xPix {
			fmt.Println("Image", i, "\t pixel", j, "\t r g b a:", pixel)
			if j == 10 {
				i++
				break
			}
		}
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func getPaths() chan string {
	ch := make(chan string)
	const dir = "../00_images/"

	go func(){
		filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			ch <- path
			return nil
		})
		close(ch)
	}()

	return ch
}

func getPixels(ch chan string) chan []pixel {
	chI := make(chan []pixel)

	go func(){
		for c := range ch {
			img := loadImage(c)
			bounds := img.Bounds()
			fmt.Println(bounds.Dx(), " x ", bounds.Dy()) // debugging
			pixels := make([]pixel, bounds.Dx()*bounds.Dy())

			for i := 0; i < bounds.Dx()*bounds.Dy(); i++ {
				x := i % bounds.Dx()
				y := i / bounds.Dx()
				r, g, b, a := img.At(x, y).RGBA()
				pixels[i].r = r
				pixels[i].g = g
				pixels[i].b = b
				pixels[i].a = a
			}
		chI <- pixels
		}
		close(chI)
	}()

	return chI
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