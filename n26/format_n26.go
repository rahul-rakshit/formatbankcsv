package n26

import (
	"encoding/csv"
	"errors"
	"fmt"
	"strings"
)

const expectedInputHeader = "\"Date\",\"Payee\",\"Account number\",\"Transaction type\",\"Payment reference\",\"Amount (EUR)\",\"Amount (Foreign Currency)\",\"Type Foreign Currency\",\"Exchange Rate\""

func FormatN26(inputLines []string) ([]string, error) {
	outputLines := []string{"\"Date\",\"Vendor\",\"Reference\",\"Amount\""}

	for index, inputLine := range inputLines {
		if index == 0 {
			if inputLine != expectedInputHeader {
				return []string{}, errors.New("Unexpected header for n26 format")
			}

			continue
		}

		csvReader := csv.NewReader(strings.NewReader(inputLine))
		row, _ := csvReader.Read()

		date := row[0]
		vendor := row[1]
		reference := row[4]
		amount := row[5]

		outputLine := fmt.Sprintf("\"%s\",\"%s\",\"%s\",\"%s\"", date, vendor, reference, amount)
		outputLines = append(outputLines, outputLine)
	}

	return outputLines, nil
}
