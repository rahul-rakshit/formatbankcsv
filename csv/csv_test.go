package csv

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCsv(t *testing.T) {
  data := [][]string{
    {"1", "2", "3"},
    {"a", "b", "c"},
  }
  filePath := ".test_temp/test.txt"

  WriteCsv(data, filePath, "^")
  fileContents, _ := ReadCsv(filePath, "^", 0)

  assert.Equal(t, data, fileContents)
}

func TestCsv_CanSkipLines(t *testing.T) {
	fileLines := []string{
		"Line 1",
		"",
		"Line 3",
		"Some header",
		"1;2;3",
		"a;b;c",
	}
	filePath := ".test_temp/complex.csv"
	file, _ := os.Create(filePath)
	for _, line := range fileLines {
		file.WriteString(line + "\n")
	}
	file.Close()

  linesToSkip := 4
	parsedData, _ := ReadCsv(filePath, ";", linesToSkip)

	expectedData := [][]string{
		{"1", "2", "3"},
		{"a", "b", "c"},
	}
	assert.Equal(t, expectedData, parsedData)
}
