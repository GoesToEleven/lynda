package main

import (
	"fmt"
	"strings"
)

func main() {
	ans := [][]string{}
	options := []string{"A", "B", "C"}
	for _, x := range options {
		for _, y := range options {
			for _, z := range options {
				if x != y && x != z && y != z {
					ans = append(ans, strings.Fields(x+" "+y+" "+z))
				}
			}
		}
	}
	for _, v := range ans {
		fmt.Println(v)
	}
}
