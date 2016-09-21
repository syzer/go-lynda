package main

import (
	"encoding/csv"
	"fmt"
	"github.com/syzer/go-lynda/lib"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://localhost:8080/Environmental_Data_Deep_Moor_2012.txt")
	if err != nil {
		log.Fatal(err)
	}
	rdr := csv.NewReader(res.Body)
	rdr.Comma = '\t'
	rdr.TrimLeadingSpace = true

	defer res.Body.Close()

	rows, err := rdr.ReadAll()
	if err != nil {
		panic(err)
	}
	fmt.Println("Total ", len(rows)-1)
	fmt.Println("Mean Air ", lib.Mean(rows, 1), " Median ", lib.Median(rows, 1))
	fmt.Println("Mean Barometric ", lib.Mean(rows, 2), " Median ", lib.Median(rows, 2))
	fmt.Println("Mean Wind Speed ", lib.Mean(rows, 7), " Median ", lib.Median(rows, 7))
}
