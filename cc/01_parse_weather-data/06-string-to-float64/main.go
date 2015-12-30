package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("../Environmental_Data_Deep_Moor_2015.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rdr := csv.NewReader(f)
	rdr.Comma = '\t'
	rdr.TrimLeadingSpace = true
	rows, err := rdr.ReadAll()
	if err != nil {
		panic(err)
	}

	for i, row := range rows {
		if i != 0 && i < 10 {
			at, _ := strconv.ParseFloat(row[1], 64)
			bm, _ := strconv.ParseFloat(row[2], 64)
			ws, _ := strconv.ParseFloat(row[7], 64)
			fmt.Printf("%T %T %T\n", at, bm, ws)
		}
	}
}
