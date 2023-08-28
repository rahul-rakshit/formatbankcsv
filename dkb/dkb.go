package dkb

import (
	"rahul-rakshit/formatbankcsv/csv"
	"rahul-rakshit/formatbankcsv/utils"
)

type Formatter struct{}

func (f Formatter) ReadCsv(inputPath string) ([][]string, error) {
	delimiter := ";"
	skipLines := 4

	return csv.ReadCsv(inputPath, delimiter, skipLines)
}

func (f Formatter) Format(csvData [][]string, fromDate string, toDate string) ([][]string, error) {
	formatted, formatErr := formatDkb(csvData)
	if formatErr != nil {
		return [][]string{}, formatErr
	}

	filtered, filterErr := utils.FilterByDate(formatted, fromDate, toDate)
	if filterErr != nil {
		return [][]string{}, filterErr
	}

	return utils.SortByDate(filtered), nil
}

func (f Formatter) WriteCsv(data [][]string, outputPath string) error {
	delimiter := ","

	return csv.WriteCsv(data, outputPath, delimiter)
}
