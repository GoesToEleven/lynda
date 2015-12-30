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

	var atTotal, bmTotal, wsTotal, counter float64

	for i, row := range rows {
		if i != 0 {
			at, _ := strconv.ParseFloat(row[1], 64)
			bm, _ := strconv.ParseFloat(row[2], 64)
			ws, _ := strconv.ParseFloat(row[7], 64)
			atTotal += at
			bmTotal += bm
			wsTotal += ws
			counter++
		}
	}
	fmt.Println("Total Records: ", counter)
	fmt.Println("Mean Air Temp: ", atTotal/counter)
	fmt.Println("Mean Barometric: ", bmTotal/counter)
	fmt.Println("Mean Wind Speed: ", wsTotal/counter)
}
