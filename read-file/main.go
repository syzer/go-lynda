package main

import (
	"encoding/csv"
	"fmt"
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

	rows, err := rdr.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, row := range rows {
		fmt.Println(row)
	}
}
