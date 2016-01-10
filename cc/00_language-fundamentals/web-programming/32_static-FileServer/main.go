package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":9000", http.FileServer(http.Dir(".")))
}

/*
All material is licensed under the Apache License Version 2.0, January 2004
http://www.apache.org/licenses/LICENSE-2.0
*/
