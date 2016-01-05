package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://lpo.dt.navy.mil/data/DM/Environmental_Data_Deep_Moor_2015.txt")
	if err != nil {
		log.Fatal(err)
	}
	bs, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", bs)
}
