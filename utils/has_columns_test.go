package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasColumns(t *testing.T) {
	headerColumn := []string{"Item", "Description", "Amount"}
	requiredColumns := []string{"Amount", "Item"}

	assert.True(t, HasColumns(headerColumn, requiredColumns))
}

func TestDoesNotHaveColumns(t *testing.T) {
	headerColumn := []string{"Item", "Description", "Amount"}
	requiredColumns := []string{"Supermarket", "Item"}

	assert.False(t, HasColumns(headerColumn, requiredColumns)) // Corrected to assert.False
}
