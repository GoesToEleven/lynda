1. research - how to read a jpeg?
	* all packages
		* https://godoc.org/?q=jpeg
	* standard library
		* https://godoc.org/image
		* https://godoc.org/image/jpeg
			* jpeg.Decode() looks promising
1. type Image
	* https://godoc.org/image#Image
		* ColorModel() color.Model
		* Bounds() Rectangle
		* At(x, y int) color.Color
1. type Rectangle
	* https://godoc.org/image#Rectangle
		* Min, Max Point
		* many methods associated with type Rectangle
			* Dx, Dy
1. type Point
	* https://godoc.org/image#Point
		* X, Y int

```
OUTPUT:

IMAGE TYPE: *image.YCbCr
Width x Height:  3867  x  2577
Total Pixels:  9965259
First pixel: 30884 27783 21662 65535
Last pixel: 0 34678 0 65535
```

1. YCbCr
	* https://godoc.org/image#YCbCr
		* implements type image.Image interface
		 	* https://godoc.org/image#Image
			* has these three methods
				* At(x, y int) color.Color
				* Bounds() Rectangle
				* ColorModel() color.Model
		* We can use
			* At(x, y int) color.Color
			* color.Color.RGBA()

1. 2^16 = 65536
	* each pixel has 4 * 16 = 64 bits
	* "Images can have 64-bit pixels with 48-bit color and a 16-bit alpha channel."
		* https://en.wikipedia.org/wiki/Color_depth
1. color depth
	* images are made up of pixels
    * pixels store color information:
		* color-depth = 1: 2 colors, monochrome
		* color-depth = 2: 4 colors
		* color-depth = 4: 16 colors
		* color-depth = 8: 256 colors
		* color-depth = 24: 16,777,216 colors, "truecolor"