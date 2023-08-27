package dkb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertDkbDate_HappyPath(t *testing.T) {
  input := "02.01.23"

  output, _ := convertDkbDate(input)
  expectedDate := "2023-01-02"

	assert.Equal(t, expectedDate, output)
}

func TestConvertDkbDate_SadPath(t *testing.T) {
  input := "some.invalid.date"

  _, err := convertDkbDate(input)
  expectedRegex := "^.*cannot parse.*$"

  assert.Regexp(t, expectedRegex, err.Error())
}
