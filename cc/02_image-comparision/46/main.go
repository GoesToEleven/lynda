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
)

// at terminal:
// go run -race main.go

// the threshold difference between pixels' rgba values
const threshold = 1000

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
	const dir = "../00_images/"
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
	name := xs[(len(xs) - 1):][0]
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

			// compareColor
			var sequence bool
			// index position for needle
			// index position for haystack
			// difference between their rgba values
			// the position of "x" when a match has been found
			// a "match" is when the diff between the pixels is less than threshold
			var diff int64
			var nIdx, hIdx, xStart int
			// compare top left pixel of needle to every pixel in haystack.
			// when the diff between the two pixels is less than threshold:
			// compare sequence of pixels in needle to sequence of pixels in haystack
			for i := 0; i < (haystack.width * haystack.height); i++ {

				// current pixel in haystack
				x := i % (haystack.width)
				y := i / (haystack.width)
//				fmt.Println(y, "\t\t", x)

				// the height of the "needle" must fit within the "haystack"
				if haystack.height-y < needle.height {
					break
				}

				// set pixels for comparison
				if !sequence {
					nIdx = 0
					hIdx = i
				}

				// current X pixel in haystack will be zero when on a newline
				if sequence && x == 0 {
					nIdx++
					// sets hIdx to starting "x" position on next row
					i += haystack.width
					hIdx = i
				}
				if sequence && x != 0 {
					nIdx++
					hIdx++
				}

				// find the diff between the pixels r,g,b,a values
				diff = compareColors(
					haystack.pixels[hIdx].r,
					haystack.pixels[hIdx].g,
					haystack.pixels[hIdx].b,
					haystack.pixels[hIdx].a,
					needle.pixels[nIdx].r,
					needle.pixels[nIdx].g,
					needle.pixels[nIdx].b,
					needle.pixels[nIdx].a)

				if diff <= threshold {
					// show what has been compared
					fmt.Println(y, "\t", x, "\t", diff)
//					fmt.Println("DIFF", diff)
//					hR := fmt.Sprintf(strconv.FormatUint(uint64(haystack.pixels[hIdx].r), 10))
//					hG := fmt.Sprintf(strconv.FormatUint(uint64(haystack.pixels[hIdx].g), 10))
//					hB := fmt.Sprintf(strconv.FormatUint(uint64(haystack.pixels[hIdx].b), 10))
//					hA := fmt.Sprintf(strconv.FormatUint(uint64(haystack.pixels[hIdx].a), 10))
//					nR := fmt.Sprintf(strconv.FormatUint(uint64(needle.pixels[nIdx].r), 10))
//					nG := fmt.Sprintf(strconv.FormatUint(uint64(needle.pixels[nIdx].g), 10))
//					nB := fmt.Sprintf(strconv.FormatUint(uint64(needle.pixels[nIdx].b), 10))
//					nA := fmt.Sprintf(strconv.FormatUint(uint64(needle.pixels[nIdx].a), 10))
//					fmt.Printf("H: INDEX - PIXEL - RGBA: %v - %v %v %v %v - (%v-%v) \n", i, hR, hG, hB, hA, x, y)
//					fmt.Printf("N: INDEX - PIXEL - RGBA: %v - %v %v %v %v \n\n", i, nR, nG, nB, nA)
					// toggle sequence
					sequence = true
					if xStart == 0 {
						xStart = x
					}
				} else {
					sequence = false
					xStart = 0
				}

			}

			// DEBUGGING - line below is for testing - needs to be changed
			matches[needle.name] = haystack.name
		}
	}
	return matches
}

func compareColors(r, g, b, a, rn, gn, bn, an uint32) int64 {
	var diff int64
	diff += int64(math.Abs(float64(r - rn)))
	diff += int64(math.Abs(float64(g - gn)))
	diff += int64(math.Abs(float64(b - bn)))
	diff += int64(math.Abs(float64(a - an)))
	return diff
}
