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
	fmt.Printf("%d %d %d %d \n", r, g, b, a) // 30884 27783 21662 65535
	fmt.Printf("%b %b %b %b \n", r, g, b, a) // 111 1000 1010 0100 110110010000111 101010010011110 1111111111111111
	fmt.Printf("%x %x %x %x \n", r, g, b, a) // 78a4 6c87 549e ffff
	fmt.Printf("%#x %#x %#x %#x \n", r, g, b, a) // 0x78a4 0x6c87 0x549e 0xffff
	fmt.Printf("%#0X %#0X %#0X %#0X \n", r, g, b, a) // 0X78A4 0X6C87 0X549E 0XFFFF
}
