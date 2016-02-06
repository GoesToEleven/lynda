package main

import (
	"fmt"
	"github.com/rwcarlsen/goexif/exif"
	"os"
)

// IMPORTANT:
// run this at terminal to import exif package:
// go get github.com/rwcarlsen/goexif/exif

func main() {
	fname := "03.jpg"

	f, _ := os.Open(fname)
	x, _ := exif.Decode(f)

	if x != nil {
		tm, _ := x.DateTime()
		fmt.Println("Taken: ", tm)

		lat, long, _ := x.LatLong()
		fmt.Println("lat, long: ", lat, ", ", long)
	}

}
