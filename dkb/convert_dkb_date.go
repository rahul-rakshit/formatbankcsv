package dkb

import "time"

func convertDkbDate(input string) (string, error) {
	inputFormat := "02.01.06"
	outputFormat := "2006-01-02"

	parsedDate, inputFormatErr := time.Parse(inputFormat, input)
	if inputFormatErr != nil {
		return "", inputFormatErr
	}

	return parsedDate.Format(outputFormat), nil
}
