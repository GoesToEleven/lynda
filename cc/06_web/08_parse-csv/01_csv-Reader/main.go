package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	// #1 open a file
	f, err := os.Open("data/first_semester.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// #2 parse a csv file
	rdr := csv.NewReader(f)
	records, err := rdr.ReadAll()
	if err != nil {
		panic(err)
	}

	// #3 display output
	for _, row := range records {
		fmt.Println(row)
	}

}
