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
	fmt.Println("Air Temp:\t", mean(rows, 1), median(rows, 1))
	fmt.Println("Barometric:\t", mean(rows, 2), median(rows, 2))
	fmt.Println("Wind Speed:\t", mean(rows, 7), median(rows, 7))
}

func median(rows [][]string, idx int) float64 {

	// to hold data which will be sorted
	var sorted []float64

	// populate the sorted []float64
	for i, row := range rows {
		if i != 0 {
			val, _ := strconv.ParseFloat(row[idx], 64)
			sorted = append(sorted, val)
		}
	}

	// Before sorting the []float64; the first 10 records.
	fmt.Println("BEFORE SORTING")
	for i, v := range sorted {
		fmt.Println(v)
		if i == 9 {
			break
		}
	}

	// Sorting the []float64
	sort.Float64s(sorted)

	// After sorting the []float64; the first 10 records.
	fmt.Println("AFTER SORTING")
	for i, v := range sorted {
		fmt.Println(v)
		if i == 9 {
			break
		}
	}
	return 0.0
}

func mean(rows [][]string, idx int) float64 {
	var total float64
	for i, row := range rows {
		if i != 0 {
			val, _ := strconv.ParseFloat(row[idx], 64)
			total += val
		}
	}
	return total / float64(len(rows)-1)
}
