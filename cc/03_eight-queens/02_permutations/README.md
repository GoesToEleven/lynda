```
	str := "ABC"
	n := len(str)
	fmt.Println(n) 			// 3

	bs := []byte(str)
	fmt.Println(bs) 		// [65 66 67]
	fmt.Printf("%T\n",bs)	// []uint8
	fmt.Println(string(bs))	// ABC

https://play.golang.org/p/c5urxfLerM
```