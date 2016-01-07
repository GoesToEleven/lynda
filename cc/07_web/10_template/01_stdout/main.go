package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
	"text/template"
)

type course struct {
	number, name, units string
}

func main() {
	// #1 get data
	semesters := academicYear("data/first_semester.txt", "data/second_semester.txt")

	// #2 parse template
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// #3 execute template
	err = tpl.Execute(os.Stdout, semesters)
	if err != nil {
		log.Fatalln(err)
	}

}

func academicYear(s ...string) [][]course {
	semesters := [][]course{}
	for _, v := range s {
		records := getRecords(v)
		semester := make([]course, 0, len(records))
		// parse data
		for i, row := range records {
			if i != 0 {
				c := course{}
				xs := strings.SplitN(row[0], " ", 2)
				c.number = xs[0]
				c.name = xs[1]
				c.units = row[1]
				semester = append(semester, c)
			}
		}
		semesters = append(semesters, semester)
	}
	return semesters
}

func getRecords(path string) [][]string {
	// open a file
	f, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// parse a csv file
	rdr := csv.NewReader(f)
	rows, err := rdr.ReadAll()
	if err != nil {
		panic(err)
	}
	return rows
}