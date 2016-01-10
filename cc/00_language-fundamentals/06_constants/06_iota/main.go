package main

import "fmt"

const (
	_ = iota      // 0
	b = iota * 10 // 1 * 10
	c = iota * 10 // 2 * 10
)

func main() {
	fmt.Println(b)
	fmt.Println(c)
}

/*
All material is licensed under the Apache License Version 2.0, January 2004
http://www.apache.org/licenses/LICENSE-2.0
*/
