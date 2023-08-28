package n26

import (
	"errors"
	"rahul-rakshit/formatbankcsv/output"
	"rahul-rakshit/formatbankcsv/utils"
)

func FormatN26(inputLines [][]string) ([][]string, error) {
	expectedInputHeader := []string{"Date", "Payee", "Account number", "Transaction type", "Payment reference", "Amount (EUR)", "Amount (Foreign Currency)", "Type Foreign Currency", "Exchange Rate"}
	outputLines := [][]string{
		output.Header,
	}

	for index, inputLine := range inputLines {
		if index == 0 {
			if !utils.IsLineEqual(inputLine, expectedInputHeader) {
				return [][]string{}, errors.New("Unexpected header for n26 format")
			}

			continue
		}

		date := inputLine[0]
		vendor := inputLine[1]
		reference := inputLine[4]
		amount := inputLine[5]

		outputLines = append(outputLines, []string{date, vendor, reference, amount})
	}

	return outputLines, nil
}
