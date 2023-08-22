package main

import (
	"fmt"
)

func handleOptions(inputPath string, outputPath string, fromDate string, toDate string, format string) error {
	fmt.Printf("Input: %s\n", inputPath)
	fmt.Printf("Output: %s\n", outputPath)
	fmt.Printf("From: %s\n", fromDate)
	fmt.Printf("To: %s\n", toDate)
	fmt.Printf("Format: %s\n", format)

  return nil
}
