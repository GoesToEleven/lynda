package main

import (
	"fmt"
	"github.com/rwcarlsen/goexif/exif"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// IMPORTANT:
// run this at terminal to import exif package:
// go get github.com/rwcarlsen/goexif/exif

func main() {

	filepath.Walk("../../", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(path)
		switch ext {
		case ".jpg", ".jpeg":
			// fmt.Println(ext)

			f, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()

			xi(f)
		}

		return nil
	})
}

func xi(f *os.File) {
	x, _ := exif.Decode(f)
	if x != nil {
		str := x.String()
		if strings.Contains(str, "ImageDescription") {

			// ImageDescription: "CROPPED and FLIPPED Statue of Liberty"
			phrase := `ImageDescription: "`
			start := strings.Index(str, phrase) + len(phrase)
			fmt.Println("DID I GET DESCRIP:", str[start:])
		}
	}
}
