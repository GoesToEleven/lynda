package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
	"text/template"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
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

func academicYear(s ...string) []semester {
	semesters := []semester{}
	for _, v := range s {
		s := semester{}
		records := getRecords(v)
		xc := make([]course, 0, len(records))
		// #3 get data
		for i, row := range records {
			if i == 0 {
				xs := strings.SplitN(row[0], ",", 2)
				t := xs[0]
				s.Term = t
			} else {
				c := course{}
				xs := strings.SplitN(row[0], " ", 2)
				c.Number = xs[0]
				c.Name = xs[1]
				c.Units = row[1]
				xc = append(xc, c)
			}
		}
		s.Courses = xc
		semesters = append(semesters, s)
	}
	return semesters
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
