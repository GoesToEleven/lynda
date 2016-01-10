package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	foo()
}

func foo() {
	// no access to x
	// this does not compile
	fmt.Println(x)
}

/* THIS WAS ADDED TO THE FILE
and this is the second line
AND THIS IS THE THIRD LINE
*/
