package main

import (
	"fmt"
	"net/http"
)

func getCookie(res http.ResponseWriter, req *http.Request) {
	cookie, _ := req.Cookie("my-cookie")
	fmt.Println(cookie)
}

func main() {
	http.HandleFunc("/", getCookie)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":9000", nil)
}

/*
All material is licensed under the Apache License Version 2.0, January 2004
http://www.apache.org/licenses/LICENSE-2.0
*/
