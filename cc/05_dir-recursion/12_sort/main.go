package main

import (
	"fmt"
	"github.com/rwcarlsen/goexif/exif"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// IMPORTANT:
// run this at terminal to import exif package:
// go get github.com/rwcarlsen/goexif/exif

func main() {

	xs := getImages()
	sort.Strings(xs)
	for i, v := range xs {
		fmt.Println(i, v)
	}

}

func getImages() []string {

	var paths []string

	filepath.Walk("../../", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(path)
		switch ext {
		case ".jpg", ".jpeg":
			f, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()

			str := xi(f)
			if str != "" {
				str = str + "|" + path
				paths = append(paths, str)
			}
		}
		return nil
	})

	return paths
}

func xi(f *os.File) string {
	x, _ := exif.Decode(f)
	if x != nil {
		str := x.String()
		if strings.Contains(str, "ImageDescription") {

			// ImageDescription: "CROPPED and FLIPPED Statue of Liberty"
			phrase := `ImageDescription: "`
			start := strings.Index(str, phrase) + len(phrase)
			end := start + strings.Index(str[start:], `"`)
			return str[start:end]
		}
	}
	return ""
}
