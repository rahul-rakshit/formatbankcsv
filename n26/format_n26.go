package n26

import (
	"encoding/csv"
	"fmt"
	"strings"
)

func FormatN26(inputLines []string) ([]string, error) {
	outputLines := []string{"\"Date\",\"Vendor\",\"Reference\",\"Amount\""}

	for index, inputLine := range inputLines {
		if index == 0 {
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
