package main

import (
	"fmt"
	stdimage "image"
	"image/jpeg"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
	_ "strconv"
	"io"
	"image/color"
	_ "image/draw"
)

// at terminal:
// go run -race main.go

const dir = "../00_images/"
// the threshold difference between pixels' rgba values
const threshold = 400

type pixel struct {
	r, g, b, a uint32
}

type image struct {
	name   string
	pixels []pixel
	width  int
	height int
}

func main() {
	start := time.Now()

	images, err := getImages()
	if err != nil {
		log.Println("Error getting images", err)
	}

	// compare the images
	matches := compare(images)

	for k, v := range matches {
		fmt.Println("This picture", k, "was in this picture", v)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func getImages() ([]image, error) {

	paths, err := getPaths()
	if err != nil {
		log.Println("Error getting paths", err)
	}

	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(len(paths))

	var images []image
	for _, path := range paths {
		go func(path string) {
			image := getPixels(path)

			mu.Lock()
			{
				images = append(images, image)
			}
			mu.Unlock()

			wg.Done()
		}(path)
	}

	wg.Wait()

	return images, nil
}

func getPaths() ([]string, error) {
	var paths []string

	wf := func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		paths = append(paths, path)
		return nil
	}

	if err := filepath.Walk(dir, wf); err != nil {
		return nil, err
	}

	return paths, nil
}

func getPixels(path string) image {
	img := loadImage(path)
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

	xs := strings.Split(path, "/")
//	name := xs[(len(xs) - 1):][0]
	name := xs[(len(xs) - 1)]
	image := image{
		name:   name,
		pixels: pixels,
		width:  bounds.Dx(),
		height: bounds.Dy(),
	}
	return image
}

func loadImage(filename string) stdimage.Image {
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

func compare(images []image) map[string]string {

	matches := make(map[string]string)

	// for each image, compare it to the other images
	for _, needle := range images {
		for _, haystack := range images {
			// needle is greater than haystack in width or height
			// needle isn't from haystack
			if needle.width > haystack.width {
				continue
			}
			if needle.height > haystack.height {
				continue
			}
			// don't compare with self
			if needle.name == haystack.name {
				continue
			}

			of, err := os.Open(dir + haystack.name)
			if err != nil {
				log.Fatalln("my program broke opening: ", err.Error())
			}
			defer of.Close()

			nf, err := os.Create("newFile.jpg")
			if err != nil {
				log.Fatalln("my program broke creating: ", err.Error())
			}
			defer nf.Close()

			io.Copy(nf, of)
			nf.Seek(0, 0)

			dnf, _ := jpeg.Decode(nf)
			b := dnf.Bounds()
			m := stdimage.NewRGBA(b)
//			draw.Draw(m, b, dnf, b.Min, draw.Src)

//			m := stdimage.NewRGBA(stdimage.Rect(0, 0, haystack.width, haystack.height))

			// compareColor
			// index position for needle
			// index position for haystack
			// difference between their rgba values
			// the position of "x" when a match has been found
			// a "match" is when the diff between the pixels is less than threshold
			var diff int64
			var nIdx, hIdx int
			// compare top left pixel of needle to every pixel in haystack.
			// when the diff between the two pixels is less than threshold:
			// compare sequence of pixels in needle to sequence of pixels in haystack
			for i := 0; i < (haystack.width * haystack.height); i++ {

				// current pixel in haystack
				x := i % (haystack.width)
				y := i / (haystack.width)

				// the height of the "needle" must fit within the "haystack"
				if haystack.height-y < needle.height {
					break
				}
				// the width of the "needle" must fit within the "haystack"
				if haystack.width-x < needle.width {
					continue
				}

				hIdx = i

				// find the diff between the pixels
				diff = compareColors(needle.pixels[nIdx], haystack.pixels[hIdx])

				if diff <= threshold {
					fmt.Println(y, "\t", x, "\t", diff)
//					d := uint8((diff/threshold)*255)
					m.Set(x, y, color.RGBA{215,12,125, 1})
					// toggle sequence
					nIdx++
					if (nIdx % needle.width) == 0 {
						i += (haystack.width - needle.width)
					}
				} else {
					nIdx = 0
				}
			}
			o := &jpeg.Options{50}
			jpeg.Encode(nf, m, o)

			// DEBUGGING - line below is for testing - needs to be changed
			matches[needle.name] = haystack.name
		}
	}
	return matches
}

func compareColors(n, h pixel) int64 {
	var diff int64
	diff += int64(math.Abs(float64(n.r - h.r)))
	diff += int64(math.Abs(float64(n.g - h.g)))
	diff += int64(math.Abs(float64(n.b - h.b)))
	diff += int64(math.Abs(float64(n.a - h.a)))
	return diff
}