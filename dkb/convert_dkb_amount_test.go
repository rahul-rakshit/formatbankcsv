package dkb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertDkbAmount_HappyPath(t *testing.T) {
	input := "-55,94 €"

	output, _ := convertDkbAmount(input)
	expectedAmount := "-55.94"

	assert.Equal(t, expectedAmount, output)
}

func TestConvertDkbAmount_SadPath(t *testing.T) {
	input := "asdf"

	_, err := convertDkbAmount(input)
	expectedError := "Failed to parse amount asdf"

	assert.Equal(t, expectedError, err.Error())
}
