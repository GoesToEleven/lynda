package main

import (
	stdimage "image"
	"image/color"
	"image/jpeg"
	"log"
	"math"
	"os"
	"strconv"
)

func mkImg(r result) {

	n := r.needle
	h := r.haystack

	of, err := os.Open(dir + h.name)
	if err != nil {
		log.Fatalln("my program broke opening: ", err.Error())
	}
	defer of.Close()

	nf, err := os.Create(n.name + "_IN_" + h.name + "_DIFFERENCE_" + strconv.Itoa(r.avgDiff) +
		"_Y_X_" + strconv.Itoa(r.hIdx/(h.width)) + "_" + strconv.Itoa(r.hIdx%(h.width)) + ".jpg")
	if err != nil {
		log.Fatalln("my program broke creating: ", err.Error())
	}
	defer nf.Close()

	dof, _ := jpeg.Decode(of)
	m := stdimage.NewRGBA(dof.Bounds())

	hIdx := r.hIdx

	for i := 0; i < n.height*n.width; i++ {

		diff := pixelDiff(n.pixels[i], h.pixels[hIdx])

		var d uint8
		a := uint8(math.Floor(255 * .5))
		if diff < threshold {
			d = uint8(255 * (1 - (diff / threshold)))
		} else {
			d = 0
		}

		// create frame
		nX := i % (n.width)
		nY := i / (n.width)
		if nX < 5 || ((n.width - nX) < 5) {
			d = 160
		}
		if nY == 0 || (nY == (n.height - 1)) {
			d = 160
		}

		hX := hIdx % (h.width)
		hY := hIdx / (h.width)
		m.Set(hX, hY, color.RGBA{d, d, d, a})

		if ((i + 1) % n.width) == 0 {
			hIdx += (h.width - n.width)
		}
		hIdx++
	}

	jpeg.Encode(nf, m, nil)
}
