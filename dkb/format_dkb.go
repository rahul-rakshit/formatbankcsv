package dkb

import (
	"errors"
	"rahul-rakshit/formatbankcsv/constants"
	"rahul-rakshit/formatbankcsv/utils"
)

func FormatDkb(inputLines [][]string) ([][]string, error) {
	expectedInputHeader := []string{"Buchungsdatum", "Wertstellung", "Status", "Zahlungspflichtige*r", "Zahlungsempfänger*in", "Verwendungszweck", "Umsatztyp", "Betrag", "Gläubiger-ID", "Mandatsreferenz", "Kundenreferenz"}
	outputLines := [][]string{
		constants.OutputHeader,
	}

	for index, inputLine := range inputLines {
		if index == 0 {
			if !utils.IsLineEqual(inputLine, expectedInputHeader) {
				return [][]string{}, errors.New("Unexpected header for dkb format")
			}

			continue
		}

		date, _ := convertDkbDate(inputLine[0])
		reference := inputLine[5]
		umsatztyp := inputLine[6]
		amount, _ := convertDkbAmount(inputLine[7])
		var vendor string

		if umsatztyp == "Ausgang" {
			vendor = inputLine[4]
		} else {
			vendor = inputLine[3]
		}

		outputLines = append(outputLines, []string{date, vendor, reference, amount})
	}

	return utils.SortByDate(outputLines), nil
}
