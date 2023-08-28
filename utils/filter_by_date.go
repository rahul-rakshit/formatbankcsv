package utils

import (
	"rahul-rakshit/formatbankcsv/output"
	"time"
)

func FilterByDate(data [][]string, fromDate string, toDate string) ([][]string, error) {
	var filteredData [][]string
	for _, row := range data {
		err := validateDates(row, fromDate, toDate)
		if err != nil {
			return [][]string{}, err
		}

		if isHeaderRow(row) || (meetsFromDateBound(row, fromDate) && meetsToDateBound(row, toDate)) {
			filteredData = append(filteredData, row)
		}
	}

	return filteredData, nil
}

func validateDates(row []string, fromDate string, toDate string) error {
	layout := "2006-01-02"

	_, rowDateErr := time.Parse(layout, row[0])
	if rowDateErr != nil && row[0] != output.Header[0] {
		return rowDateErr
	}

	_, fromDateErr := time.Parse(layout, fromDate)
	if fromDateErr != nil && fromDate != "" {
		return fromDateErr
	}

	_, toDateErr := time.Parse(layout, fromDate)
	if toDateErr != nil && toDate != "" {
		return fromDateErr
	}

	return nil
}

func isHeaderRow(row []string) bool {
	return row[0] == output.Header[0]
}

func meetsFromDateBound(row []string, fromDateString string) bool {
	if fromDateString == "" {
		return true
	}

	layout := "2006-01-02"
	rowDate, _ := time.Parse(layout, row[0])
	fromDate, _ := time.Parse(layout, fromDateString)

	return rowDate.After(fromDate) || rowDate.Equal(fromDate)
}

func meetsToDateBound(row []string, toDateString string) bool {
	if toDateString == "" {
		return true
	}

	layout := "2006-01-02"
	rowDate, _ := time.Parse(layout, row[0])
	toDate, _ := time.Parse(layout, toDateString)

	return rowDate.Before(toDate) || rowDate.Equal(toDate)
}
