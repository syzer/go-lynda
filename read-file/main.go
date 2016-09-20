package main

import (
	"encoding/csv"
	"fmt"
	"github.com/syzer/go-lynda/lib"
	"os"
)

func main() {
	f, err := os.Open("./LPO_weatherdata/Environmental_Data_Deep_Moor_2012.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rdr := csv.NewReader(f)
	rdr.Comma = '\t'
	rdr.TrimLeadingSpace = true

	rows, err := rdr.ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Println("Total ", len(rows)-1)
	fmt.Println("Mean Air ", lib.Mean(rows, 1), " Median ", lib.Median(rows, 1))
	fmt.Println("Mean Barometric ", lib.Mean(rows, 2), " Median ", lib.Median(rows, 2))
	fmt.Println("Mean Wind Speed ", lib.Mean(rows, 7), " Median ", lib.Median(rows, 7))
}
