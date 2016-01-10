package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := `
		["one phrase", "another phrase", "four score", "monday night"]
		`
	fmt.Printf("%T\n", jsonData)

	var data []string
	json.Unmarshal([]byte(jsonData), &data)
	fmt.Println("unmarshalled: ", data)
	fmt.Println("unmarshalled: ", data[1])
	for i, v := range data {
		fmt.Println("unmarshalled: ", i, " - ", v)
	}

	data = append(data, "One love")
	bs, _ := json.Marshal(data)
	fmt.Println("marshalled: ", string(bs))
}

/*
All material is licensed under the Apache License Version 2.0, January 2004
http://www.apache.org/licenses/LICENSE-2.0
*/
