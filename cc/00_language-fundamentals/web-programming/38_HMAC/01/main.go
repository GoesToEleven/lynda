package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	sessionID := "1123456789"
	fmt.Println(sessionID)
	code := getCode(sessionID)
	fmt.Println(code)
	// we get the same code here b/c the string is unchanged
	code = getCode(sessionID)
	fmt.Println(code)
	// we get a different code here b/c the string is unchanged
	code = getCode(sessionID + "some other stuff")
	fmt.Println(code)

}

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("this can be anything we want it to be it is our private key"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

/*
All material is licensed under the Apache License Version 2.0, January 2004
http://www.apache.org/licenses/LICENSE-2.0
*/
