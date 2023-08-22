package n26

import (
	"rahul-rakshit/formatbankcsv/csv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatN26_HappyPath(t *testing.T) {
	sampleN26Data, _ := csv.ReadCsv("./fixtures/sample_n26.csv")

	formatted, _ := FormatN26(sampleN26Data)

	expectedResultCsv, _ := csv.ReadCsv("./fixtures/expected_result.csv")
	assert.Equal(t, expectedResultCsv, formatted)
}

func TestFormatN26_BadHeaderError(t *testing.T) {
  badHeaderData, _ := csv.ReadCsv("./fixtures/bad_header_n26.csv")

  _, err := FormatN26(badHeaderData)

  assert.Equal(t, "Unexpected header for n26 format", err.Error())
}
