package main

import (
	// "fmt"
	"rahul-rakshit/formatbankcsv/csv"
	"rahul-rakshit/formatbankcsv/n26"
)

func handleOptions(inputPath string, outputPath string, fromDate string, toDate string, format string) error {
	// fmt.Printf("Input: %s\n", inputPath)
	// fmt.Printf("Output: %s\n", outputPath)
	// fmt.Printf("From: %s\n", fromDate)
	// fmt.Printf("To: %s\n", toDate)
	// fmt.Printf("Format: %s\n", format)

	data, readErr := csv.ReadCsv(inputPath, ",", 0)
	if readErr != nil { return readErr }

	formatted, formattingErr := n26.FormatN26(data)
	if formattingErr != nil { return formattingErr }

	writeErr := csv.WriteCsv(formatted, outputPath, ",")
	if writeErr != nil { return writeErr }

	return nil
}
