package csv

import (
	"encoding/csv"
	"os"
)

func ReadCsv(inputPath string) ([][]string, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return [][]string{}, err
	}
  defer file.Close()

	reader := csv.NewReader(file)

	rows, err := reader.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

  return rows, nil
}
