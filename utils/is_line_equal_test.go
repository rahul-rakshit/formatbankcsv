package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLineEqual_EqualCase(t *testing.T) {
	line := []string{"1", "2", "3"}

	assert.True(t, IsLineEqual(line, line))
}

func TestIsEqual_LengthMismatch(t *testing.T) {
	line1 := []string{"1", "2", "3"}
	line2 := []string{"1", "2", "3", "4"}

	assert.False(t, IsLineEqual(line1, line2))
}

func TestIsEqual_UnequalCase(t *testing.T) {
	line1 := []string{"1", "3", "2"}
	line2 := []string{"1", "2", "3"}

	assert.False(t, IsLineEqual(line1, line2))
}
