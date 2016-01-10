package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type course struct {
	Number, Name, Units string
}

func main() {
	records := getRecords("data/first_semester.txt")

	// #3 display output
	for i, row := range records {
		if i == 0 {
			xs := strings.SplitN(row[0], ",", 2)
			t := xs[0]
			fmt.Println("TERM:", t)
		} else {
			c := course{}
			xs := strings.SplitN(row[0], " ", 2)
			c.Number = xs[0]
			c.Name = xs[1]
			c.Units = row[1]
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
