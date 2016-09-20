package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
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

	// atmospheric, baromertic, wind speed
	var atTotal, bmTotal, wsTotal, count float64

	for i, row := range rows {
		if i != 0 {
			at, _ := strconv.ParseFloat(row[1], 64)
			bm, _ := strconv.ParseFloat(row[2], 64)
			ws, _ := strconv.ParseFloat(row[7], 64)
			atTotal += at
			bmTotal += bm
			wsTotal += ws
			count++
		}
	}
	fmt.Println("Total ", count)
	fmt.Println("Mean Air ", atTotal / count)
	fmt.Println("Mean Barometric ", bmTotal / count)
	fmt.Println("Mean  Wind Speed ", wsTotal / count)
}
