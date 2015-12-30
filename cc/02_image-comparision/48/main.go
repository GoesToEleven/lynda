package main
import (
	"image/jpeg"
	"io"
	"log"
	"os"
	"fmt"
	"image/color"
)

func main() {
	of, err := os.Open("../00_images/78771293.jpg")
	if err != nil {
		log.Fatalln("my program broke opening: ", err.Error())
	}
	defer of.Close()

	nf, err := os.Create("newFile.jpg")
	if err != nil {
		log.Fatalln("my program broke creating: ", err.Error())
	}
	defer nf.Close()

	io.Copy(nf, of)
	nf.Seek(0, 0)

	m, _ := jpeg.Decode(nf)
	fmt.Printf("%T \n",m)
}
