package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
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

	// Don't count the header row in len(rows)
	fmt.Println("Total Records: ", len(rows)-1)
	fmt.Println("Air Temp:\t", median(rows, 1))
	fmt.Println("Barometric:\t", median(rows, 2))
	fmt.Println("Wind Speed:\t", median(rows, 7))
}

func median(rows [][]string, idx int) float64 {
	var sorted []float64
	for i, row := range rows {
		if i != 0 {
			val, _ := strconv.ParseFloat(row[idx], 64)
			sorted = append(sorted, val)
		}
	}
	sort.Float64s(sorted)

	if len(sorted)%2 == 0 {
		// even number of items
		// for example: 3, 5, 8, 9
		// median is (5 + 8) / 2 = 6.5
		return 0.0
	}
	// odd number of items
	// for example: 3, 5, 8
	// median is 5
	return 1.0
}
