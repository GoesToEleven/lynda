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
	"image/color"
	_ "image/draw"
)

// at terminal:
// go run -race main.go

const dir = "../00_images/"
// the threshold difference between pixels' rgba values

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

			dof, _ := jpeg.Decode(of)
			m := stdimage.NewRGBA(dof.Bounds())

			// compareColor
			var nIdx int
			var diff int64
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

				// find the diff between the pixels
				diff = compareColors(needle.pixels[nIdx], haystack.pixels[i])
//				if diff < 1000 {
//					m.Set(x, y, color.RGBA{255,182,193,1})
//				} else {
					d := uint8(255*(diff/(255*4)))
					a := uint8(1*(diff/(255*4)))
					m.Set(x, y, color.RGBA{d,d,d,a})
//				}
			}

			nf, err := os.Create("newFile.jpg")
			if err != nil {
				log.Fatalln("my program broke creating: ", err.Error())
			}
			defer nf.Close()

			jpeg.Encode(nf, m, nil)

			// DEBUGGING - line below is for testing - needs to be changed
			matches[needle.name] = haystack.name
		}
	}
	return matches
}

func compareColors(n, h pixel) int64 {
//	fmt.Println(n.r, n.g, n.b, n.a)
	var diff int64
	diff += int64(math.Abs(float64(n.r - h.r)))
	diff += int64(math.Abs(float64(n.g - h.g)))
	diff += int64(math.Abs(float64(n.b - h.b)))
	diff += int64(math.Abs(float64(n.a - h.a)))
	return diff
}