package dkb

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func convertDkbAmount(input string) (string, error) {
	cleanedInput := strings.Replace(strings.TrimSuffix(input, " €"), ",", ".", -1)

	value, err := strconv.ParseFloat(cleanedInput, 64)
	if err != nil {
		return "", errors.New("Failed to parse amount " + input)
	}

	return fmt.Sprintf("%.2f", value), nil
}
