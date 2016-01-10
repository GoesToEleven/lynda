package main

import "fmt"

func main() {
	for i := 60; i < 122; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}

/* THIS WAS ADDED TO THE FILE
and this is the second line
AND THIS IS THE THIRD LINE
*/
