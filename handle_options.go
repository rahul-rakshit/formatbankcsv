package main

import (
	"errors"
	"rahul-rakshit/formatbankcsv/dkb"
	"rahul-rakshit/formatbankcsv/n26"
	"strings"
)

func handleOptions(inputPath string, outputPath string, fromDate string, toDate string, format string) error {
	lowerCaseFormat := strings.ToLower(format)
	var formatter CsvFormatter
	if lowerCaseFormat == "n26" {
		formatter = &n26.Formatter{}
	} else if lowerCaseFormat == "dkb" {
		formatter = &dkb.Formatter{}
	} else {
		return errors.New("Format must be either \"dkb\" or \"n26\"")
	}

	csvData, readErr := formatter.ReadCsv(inputPath)
	if readErr != nil {
		return readErr
	}

	formatted, formatErr := formatter.Format(csvData, fromDate, toDate)
	if formatErr != nil {
		return formatErr
	}

	return formatter.WriteCsv(formatted, outputPath)
}
