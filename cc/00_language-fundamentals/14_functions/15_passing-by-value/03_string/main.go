package main

import "fmt"

func main() {

	name := "Todd"
	fmt.Println(name) // Todd

	changeMe(name)

	fmt.Println(name) // Todd
}

func changeMe(z string) {
	fmt.Println(z) // Todd
	z = "Rocky"
	fmt.Println(z) // Rocky
}

/*
All material is licensed under the Apache License Version 2.0, January 2004
http://www.apache.org/licenses/LICENSE-2.0
*/
