package main

import "fmt"

func main() {
	age := 44
	changeMe(age)
	fmt.Println(age) // 44
}

func changeMe(z int) {
	fmt.Println(z)
	z = 24
}

// when changeMe is called on line 8
// the value 44 is being passed as an argument

/*
All material is licensed under the Apache License Version 2.0, January 2004
http://www.apache.org/licenses/LICENSE-2.0
*/
