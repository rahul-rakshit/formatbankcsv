package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterByDate_BoundsAreInclusive(t *testing.T) {
	data := [][]string{
		{"Date", "Name", "Amount"},
		{"1999-03-03", "Frederick", "17"},
		{"2000-07-18", "Howard", "19"},
		{"2004-02-28", "Reginald", "91"},
		{"2004-04-13", "Eduard", "95"},
		{"2006-05-15", "Arthur", "98"},
	}

	filteredData, err := FilterByDate(data, "2000-07-18", "2004-04-13")

	expectedData := [][]string{
		{"Date", "Name", "Amount"},
		{"2000-07-18", "Howard", "19"},
		{"2004-02-28", "Reginald", "91"},
		{"2004-04-13", "Eduard", "95"},
	}
	assert.Equal(t, expectedData, filteredData)
	assert.Nil(t, err)
}

func TestFilterByDate_BoundsAreOptional(t *testing.T) {
	originalData := [][]string{
		{"Date", "Name", "Amount"},
		{"2000-07-18", "Gertrud", "19"},
		{"2004-02-28", "Hedwig", "91"},
		{"2004-04-13", "Brunhilde", "95"},
	}

	filteredData, err := FilterByDate(originalData, "", "")

	assert.Equal(t, originalData, filteredData)
	assert.Nil(t, err)
}

func TestFilterByDate_DateValidation(t *testing.T) {
	originalData := [][]string{
		{"Date", "Name", "Amount"},
		{"2000-07-18", "Geneviève", "102"},
		{"123-9-mauvais", "Colette", "98"},
		{"2004-04-13", "Éléonore", "97"},
	}

	filteredData, err := FilterByDate(originalData, "", "")

	assert.Regexp(t, "^parsing time \"123-9-mauvais\".*cannot parse.*$", err.Error())
	assert.Equal(t, [][]string{}, filteredData)

}
