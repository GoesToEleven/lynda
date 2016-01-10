package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, 5*5)
	if err != nil {
		log.Fatalln(err)
	}
}

/*
In software engineering, a pipeline consists of a chain of processing elements
(processes, threads, coroutines, functions, etc.), arranged so that
the output of each element is the input of the next;
*/

/*
All material is licensed under the Apache License Version 2.0, January 2004
http://www.apache.org/licenses/LICENSE-2.0
*/
