package main

type CsvFormatter interface {
	ReadCsv(inputPath string) ([][]string, error)
	Format(csvData [][]string, fromDate string, toDate string) ([][]string, error)
	WriteCsv(data [][]string, outputPath string) error
}
