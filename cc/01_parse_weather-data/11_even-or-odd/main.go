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
	fmt.Sprintln("Total Records: ", len(rows)-1)
	fmt.Sprintln("Air Temp:\t", mean(rows, 1), median(rows, 1))
	fmt.Sprintln("Barometric:\t", mean(rows, 2), median(rows, 2))
	fmt.Sprintln("Wind Speed:\t", mean(rows, 7), median(rows, 7))
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
	fmt.Println("len(sorted): ", len(sorted))
	fmt.Println("len(sorted)/2: ", len(sorted)/2)
	fmt.Println("float64(len(sorted))/2: ", float64(len(sorted))/2)

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
