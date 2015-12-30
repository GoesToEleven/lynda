package main

import (
	"fmt"
	"github.com/disintegration/imaging"
	"os"
	"runtime"
)

func main() {

	// maximize CPU usage for maximum performance
	runtime.GOMAXPROCS(runtime.NumCPU())

	// load original image
	img, err := imaging.Open("../00_images/78771293a.jpg")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// grayscale the image
	grayimg := imaging.Grayscale(img)

	// save grayscaled image
	err = imaging.Save(grayimg, "./result.png")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// everything ok
	fmt.Println("Done")
}