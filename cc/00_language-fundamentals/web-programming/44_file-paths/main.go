package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	wd, _ := os.Getwd()
	filename := "01.jpg"
	path := filepath.Join(wd, "assets", "imgs", filename)
	fmt.Println(path)
}

/*
All material is licensed under the Apache License Version 2.0, January 2004
http://www.apache.org/licenses/LICENSE-2.0
*/
