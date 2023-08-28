package utils

import (
	"rahul-rakshit/formatbankcsv/output"
	"sort"
	"time"
)

func SortByDate(input [][]string) [][]string {
	type DateValuePair struct {
		Date  string
		Value []string
	}

	var data []DateValuePair

	for _, item := range input {
		data = append(data, DateValuePair{
			Date:  item[0],
			Value: item[1:],
		})
	}

	sort.Slice(data, func(i, j int) bool {
		if data[i].Date == output.Header[0] {
			return true
		} else if data[j].Date == output.Header[0] {
			return false
		}

		date1, err1 := time.Parse("2006-01-02", data[i].Date)
		date2, err2 := time.Parse("2006-01-02", data[j].Date)

		if err1 != nil || err2 != nil {
			return false
		}

		return date1.Before(date2)
	})

	var sorted [][]string

	for _, item := range data {
		sortedItem := append([]string{item.Date}, item.Value...)
		sorted = append(sorted, sortedItem)
	}

	return sorted
}
