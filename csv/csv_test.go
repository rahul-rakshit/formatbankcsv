package csv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCsv(t *testing.T) {
  data := [][]string{
    {"1", "2", "3"},
    {"a", "b", "c"},
  }
  filePath := ".test_temp/test.txt"


  WriteCsv(data, filePath)
  fileContents, _ := ReadCsv(filePath)

  assert.Equal(t, data, fileContents)
}
