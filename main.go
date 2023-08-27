package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	inputPath  string
	outputPath string
	fromDate   string
	toDate     string
	format     string
)

var rootCmd = &cobra.Command{
	Use:   "formatbankcsv",
	Short: "A command-line tool to format your bank's csv to Rana format",
	Run:   runMainCommand,
}

func runMainCommand(cmd *cobra.Command, args []string) {
	err := handleOptions(inputPath, outputPath, fromDate, toDate, format)

	if err != nil {
    fmt.Println(err.Error())
		os.Exit(2)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&inputPath, "input", "i", "", "Path to the input file (mandatory)")
	rootCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Path to the output file (mandatory)")
	rootCmd.Flags().StringVarP(&format, "format", "F", "", "Format (\"dkb\" or \"n26\") (mandatory)")
	rootCmd.Flags().StringVarP(&fromDate, "from", "f", "", "Start date in yyyy-mm-dd format, inclusive (optional)")
	rootCmd.Flags().StringVarP(&toDate, "to", "t", "", "End date in yyyy-mm-dd format, inclusive (optional)")
	rootCmd.MarkFlagRequired("input")
	rootCmd.MarkFlagRequired("output")
	rootCmd.MarkFlagRequired("format")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
