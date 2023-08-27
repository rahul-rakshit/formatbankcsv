package dkb

import (
  "rahul-rakshit/formatbankcsv/csv"
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestFormatDkb_HappyPath(t *testing.T) {
  sampleDkbData, _ := csv.ReadCsv("./fixtures/sample_dkb.csv", ";", 4)

  formatted, _ := FormatDkb(sampleDkbData)

  expectedResultCsv, _ := csv.ReadCsv("./fixtures/expected_result.csv", ",", 0)
  assert.Equal(t, expectedResultCsv, formatted)
}
