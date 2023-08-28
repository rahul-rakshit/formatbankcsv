package dkb

import (
	"rahul-rakshit/formatbankcsv/csv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatDkb_HappyPath(t *testing.T) {
	sampleDkbData, _ := csv.ReadCsv("./fixtures/sample_dkb.csv", ";", 4)

	formatted, _ := formatDkb(sampleDkbData)

	expectedResultCsv, _ := csv.ReadCsv("./fixtures/expected_result.csv", ",", 0)
	assert.Equal(t, expectedResultCsv, formatted)
}

func TestFormatN26_BadHeaderError(t *testing.T) {
	badHeaderData, _ := csv.ReadCsv("./fixtures/bad_header_dkb.csv", ";", 4)

	_, err := formatDkb(badHeaderData)

	assert.Equal(t, "Unexpected header for dkb format", err.Error())
}

func TestFormatN26_OldFormatError(t *testing.T) {
	_, err := csv.ReadCsv("./fixtures/old_dkb_format.csv", ";", 4)

	assert.Regexp(t, "record on line 3: wrong number of fields", err.Error())
}
