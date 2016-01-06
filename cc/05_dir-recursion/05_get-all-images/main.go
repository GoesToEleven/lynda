package main

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
)

func main() {

	filepath.Walk("../../", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(path)
		switch ext {
		case ".jpg", ".jpeg":
			fmt.Println(ext)

			f, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()

			img, err := jpeg.Decode(f)
			r, g, b, a := img.At(0, 0).RGBA()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%d %d %d %d \n", r, g, b, a) // 30884 27783 21662 65535

		}

		return nil
	})
}
