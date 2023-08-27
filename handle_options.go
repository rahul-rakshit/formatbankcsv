package main

import (
	"errors"
	"rahul-rakshit/formatbankcsv/csv"
	"rahul-rakshit/formatbankcsv/dkb"
	"rahul-rakshit/formatbankcsv/n26"
	"rahul-rakshit/formatbankcsv/utils"
	"strings"
)

func handleOptions(inputPath string, outputPath string, fromDate string, toDate string, format string) error {
	lowerCaseFormat := strings.ToLower(format)
	if lowerCaseFormat != "dkb" && lowerCaseFormat != "n26" {
		return errors.New("Format must be either \"dkb\" or \"n26\"")
	}

	var skipLines int
  var delimiter string
	if lowerCaseFormat == "n26" {
		skipLines = n26.LinesToSkip
    delimiter = n26.Delimiter
	} else {
		skipLines = dkb.LinesToSkip
    delimiter = dkb.Delimiter
	}
	originalCsvData, readErr := csv.ReadCsv(inputPath, delimiter, skipLines)
	if readErr != nil {
		return readErr
	}

	var formatted [][]string
	var formattingErr error
	if lowerCaseFormat == "n26" {
		formatted, formattingErr = n26.FormatN26(originalCsvData)
	} else {
		formatted, formattingErr = dkb.FormatDkb(originalCsvData)
	}
	if formattingErr != nil {
		return formattingErr
	}

	filtered, filterErr := utils.FilterByDate(formatted, fromDate, toDate)
	if filterErr != nil {
		return filterErr
	}

	sorted := utils.SortByDate(filtered)

	writeErr := csv.WriteCsv(sorted, outputPath, ",")
	if writeErr != nil {
		return writeErr
	}

	return nil
}
