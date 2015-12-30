package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("../00_images/78771293.jpg")
	if err != nil {
		log.Fatal(err)
	}

	fi, _ := f.Stat()

	fmt.Println("Name: \t\t", fi.Name())
	fmt.Println("Size: \t\t", fi.Size())
	fmt.Println("Mode: \t\t", fi.Mode())
	fmt.Println("ModTime: \t", fi.ModTime())
	fmt.Println("IsDir: \t\t", fi.IsDir())
	fmt.Println("Sys: \t\t", fi.Sys())

}
