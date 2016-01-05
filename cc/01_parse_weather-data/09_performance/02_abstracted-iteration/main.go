package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	f, err := os.Open("../../Environmental_Data_Deep_Moor_2015.txt")
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

	start := time.Now()
	fmt.Println("Total Records: ", len(rows)-1)
	fmt.Println("Mean Air Temp: ", mean(rows, 1))
	fmt.Println("Mean Barometric: ", mean(rows, 2))
	fmt.Println("Mean Wind Speed: ", mean(rows, 7))
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("Abstracted: %s\n", delta)
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
