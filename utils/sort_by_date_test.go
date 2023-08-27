package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortByDate(t *testing.T) {
	input := [][]string{
    {"date", "letter", "number"},
		{"1983-01-07", "a", "1"},
		{"2007-12-19", "b", "2"},
		{"1998-08-27", "c", "3"},
	}

	sorted := SortByDate(input)
	expected := [][]string{
    {"date", "letter", "number"},
		{"1983-01-07", "a", "1"},
		{"1998-08-27", "c", "3"},
		{"2007-12-19", "b", "2"},
	}
	assert.Equal(t, expected, sorted)
}
