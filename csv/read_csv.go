package csv

import (
	"bufio"
	"encoding/csv"
	"os"
	"strings"
)

func ReadCsv(inputPath string, delimiter string, skipLines int) ([][]string, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return [][]string{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for i := 0; i < skipLines; i++ {
		scanner.Scan()
	}

	remainingContent := ""
	for scanner.Scan() {
		remainingContent += scanner.Text() + "\n"
	}

	reader := csv.NewReader(strings.NewReader(remainingContent))
	reader.Comma = rune(delimiter[0])

	rows, err := reader.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return rows, nil
}
