package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 12
	y := "I have this many: " + strconv.Itoa(x)
	fmt.Println(y)
	//	fmt.Println("I have this many: ", strconv.Itoa(x), x)
}

/*
All material is licensed under the Apache License Version 2.0, January 2004
http://www.apache.org/licenses/LICENSE-2.0
*/
