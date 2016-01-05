package main

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"
)

func main() {

	f, err := os.Open("../00_images/78771293.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, err := jpeg.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	r, g, b, a := img.At(0, 0).RGBA()
	fmt.Printf("%T %T %T %T \n", r, g, b, a) // uint32 uint32 uint32 uint32
	fmt.Printf("%v %v %v %v \n", r, g, b, a) // 30884 27783 21662 65535
}
