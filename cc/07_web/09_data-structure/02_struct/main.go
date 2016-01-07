package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type course struct {
	number, name, units string
}

func main() {

	records := getRecords("data/first_semester.txt")

	// #3 display output
	for i, row := range records {
		if i != 0 {
			c := course{}
			xs := strings.SplitN(row[0], " ", 2)
			c.number = xs[0]
			c.name = xs[1]
			c.units = row[1]
			fmt.Println(c)
		}
	}

}

func getRecords(path string) [][]string {
	// #1 open a file
	f, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// #2 parse a csv file
	rdr := csv.NewReader(f)
	rows, err := rdr.ReadAll()
	if err != nil {
		panic(err)
	}
	return rows
}