package main

import (
	"fmt"
	"github.com/goestoeleven/GolangTraining/04_scope/01_package-scope/02_visibility/vis"
)

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
}

/* THIS WAS ADDED TO THE FILE
and this is the second line
AND THIS IS THE THIRD LINE
*/
