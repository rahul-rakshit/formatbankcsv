package n26

import (
	"errors"
	"rahul-rakshit/formatbankcsv/output"
	"rahul-rakshit/formatbankcsv/utils"
)

func FormatN26(inputLines [][]string) ([][]string, error) {
	requiredColumns := []string{"Value Date", "Partner Name", "Payment Reference", "Amount (EUR)"}

	header := inputLines[0]
	if !utils.HasColumns(header, requiredColumns) {
		return nil, errors.New("Unexpected header for n26 format")
	}

	columnIndices := make(map[string]int)
	for i, col := range header {
		columnIndices[col] = i
	}

	var requiredIndices []int
	for _, col := range requiredColumns {
		for i, headerCol := range header {
			if headerCol == col {
				requiredIndices = append(requiredIndices, i)
				break
			}
		}
	}

	outputLines := [][]string{
		output.Header,
	}
	for _, row := range inputLines[1:] {
		extractedRow := make([]string, len(requiredIndices))
		for i, index := range requiredIndices {
			extractedRow[i] = row[index]
		}
		outputLines = append(outputLines, extractedRow)
	}

	return outputLines, nil
}
