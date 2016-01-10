package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x = "12"
	var y = 6
	z, _ := strconv.Atoi(x)
	fmt.Println(y + z)
}

/*
All material is licensed under the Apache License Version 2.0, January 2004
http://www.apache.org/licenses/LICENSE-2.0
*/
