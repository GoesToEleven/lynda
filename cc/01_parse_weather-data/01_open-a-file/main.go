package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("../Environmental_Data_Deep_Moor_2015.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Println(f)
}
