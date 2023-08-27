package dkb

import (
	"errors"
	"rahul-rakshit/formatbankcsv/constants"
	"rahul-rakshit/formatbankcsv/csv"
)

func FormatDkb(inputLines [][]string) ([][]string, error) {
	expectedInputHeader := []string{"Buchungsdatum", "Wertstellung", "Status", "Zahlungspflichtige*r", "Zahlungsempfänger*in", "Verwendungszweck", "Umsatztyp", "Betrag", "Gläubiger-ID", "Mandatsreferenz", "Kundenreferenz"}
	outputLines := [][]string{
		constants.OutputHeader,
	}

	for index, inputLine := range inputLines {
		if index < 4 {
			continue
		}
		if index == 4 {
			if !csv.IsLineEqual(inputLine, expectedInputHeader) {
				return [][]string{}, errors.New("Unexpected header for dkb format")
			}

			continue
		}
	}

	return outputLines, nil
}
