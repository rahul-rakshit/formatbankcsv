package dkb

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func convertDkbAmount(input string) (string, error) {
	var builder strings.Builder

	for _, char := range input {
		if unicode.IsDigit(char) || char == rune("-"[0]) {
			builder.WriteRune(char)
		} else if char == rune(","[0]) {
			builder.WriteRune(rune("."[0]))
		}
	}

	value, err := strconv.ParseFloat(builder.String(), 64)
	if err != nil {
		return "", errors.New("Failed to parse amount " + input)
	}

	return fmt.Sprintf("%.2f", value), nil
}
