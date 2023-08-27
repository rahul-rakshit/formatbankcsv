package dkb

import (
	"fmt"
	"rahul-rakshit/formatbankcsv/csv"
	"testing"

	// "github.com/stretchr/testify/assert"
)

// func TestFormatN26_HappyPath(t *testing.T) {
//   sampleN26Data, _ := csv.ReadCsv("./fixtures/sample_dkb.csv")

//   formatted, _ := FormatDkb(sampleN26Data)

//   expectedResultCsv, _ := csv.ReadCsv("./fixtures/expected_result.csv")
//   assert.Equal(t, expectedResultCsv, formatted)
// }

func TestFormatN26_BadHeaderError(t *testing.T) {
	badHeaderData, _ := csv.ReadCsv("./fixtures/bad_header_n26.csv", ";")

  fmt.Println(badHeaderData)

	// data, err := FormatDkb(badHeaderData)

	// fmt.Println(data)

	// assert.Equal(t, "Unexpected header for dkb format", err.Error())
}
