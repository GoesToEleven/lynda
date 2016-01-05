package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
)

func main() {

	images := getImages()

	for _, v := range images {
		r, g, b, a := v.At(0, 0).RGBA()
		fmt.Printf("%d %d %d %d \n", r, g, b, a) // 30884 27783 21662 65535
	}

}

func getImages() []image.Image {

	var images []image.Image

	filepath.Walk("../00_images/", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		img := loadImage(path)
		images = append(images, img)
		return nil
	})

	return images
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
