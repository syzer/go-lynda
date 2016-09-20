package lib

import (
	"sort"
	"strconv"
)

func Median(rows [][]string, index int) float64 {
	var sorted []float64

	for i, row := range rows {
		if i != 0 {
			val, _ := strconv.ParseFloat(row[index], 64)
			sorted = append(sorted, val)
		}
	}
	sort.Float64s(sorted)

	// even
	if len(sorted)%2 == 0 {
		middle := len(sorted) / 2
		highElem := sorted[middle]
		lowerElem := sorted[middle-1]
		return (highElem - lowerElem) / 2
	}
	middle := len(sorted) / 2
	return sorted[middle]
}

func Mean(rows [][]string, index int) float64 {
	var total float64
	for i, row := range rows {
		if i != 0 {
			val, _ := strconv.ParseFloat(row[index], 64)
			total += val
		}
	}
	return total / float64(len(rows)-1)
}
