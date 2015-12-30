package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("../Environmental_Data_Deep_Moor_2015.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rdr := csv.NewReader(f)
	rdr.Comma = '\t'
	fmt.Println(rdr.TrimLeadingSpace)
	rdr.TrimLeadingSpace = true
	fmt.Println(rdr.TrimLeadingSpace)
	rows, err := rdr.ReadAll()
	if err != nil {
		panic(err)
	}

	for i, row := range rows {
		fmt.Println(row)
		if i == 1 {
			fmt.Printf("%T %T %T\n", row[1], row[2], row[7])
			fmt.Println(row[1], row[2], row[7])
			break
		}
	}
}
