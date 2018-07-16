package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// This url used to be at: http://lpo.dt.navy.mil/data/DM/Environmental_Data_Deep_Moor_2015.txt
	// but has since been discontinued, the below url points to the source repository
	res, err := http.Get("https://raw.githubusercontent.com/GoesToEleven/lynda/master/cc/01_parse_weather-data/Environmental_Data_Deep_Moor_2015.txt")
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
