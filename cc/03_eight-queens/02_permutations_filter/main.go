package main

import "fmt"

func main() {
	options := []string{"A", "B", "C"}
	for _, x := range options {
		for _, y := range options {
			for _, z := range options {
				if x != y && x != z && y != z {
					fmt.Println(x, y, z)
				}
			}
		}
	}
}
