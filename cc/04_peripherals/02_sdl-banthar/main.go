package main

import (
	"github.com/tanema/amore/mouse"
	"fmt"
)

func main() {
//	for {
//		x := mouse.GetX()
//		y := mouse.GetY()
//		fmt.Println(x, y)
//	}
	
	x, y := mouse.GetPosition()
	fmt.Println(x, y)

}
