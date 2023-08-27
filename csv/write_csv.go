package csv

import (
	"encoding/csv"
	"os"
)

func WriteCsv(data [][]string, outputPath string, delimiter string) error {
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
  writer.Comma = rune(delimiter[0])
	defer writer.Flush()

	for _, row := range data {
		err := writer.Write(row)
		if err != nil {
			return err
		}
	}

	return nil
}
