package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	records := getRecords("data/first_semester.txt")

	// #3 display output
	for i, row := range records {
		if i == 0 {
			xs := strings.SplitN(row[0], ",", 2)
			t := xs[0]
			fmt.Println("TERM:", t)
		} else {
			xs := strings.SplitN(row[0], " ", 2)
			cNum := xs[0]
			cName := xs[1]
			cUnits := row[1]
			fmt.Println(cNum, cName, cUnits)
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
