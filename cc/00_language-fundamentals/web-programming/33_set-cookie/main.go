package main

import (
	"net/http"
)

func setCookie(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "my-cookie",
		Value: "COOKIE MONSTER",
	})
}

func main() {
	http.HandleFunc("/", setCookie)
	http.ListenAndServe(":9000", nil)
}

/*
All material is licensed under the Apache License Version 2.0, January 2004
http://www.apache.org/licenses/LICENSE-2.0
*/
