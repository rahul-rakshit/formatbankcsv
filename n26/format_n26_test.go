package n26

import (
	"os"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("formatN26", func() {
	Context("In the happy path", func() {
		It("it deletes a few columns from the input", func() {
			sampleN26Csv, _ := os.ReadFile("./fixtures/sample_n26.csv")
			sampleN26Lines := strings.Split(string(sampleN26Csv), "\n")

			formatted, _ := FormatN26(sampleN26Lines)

			expectedResultCsv, _ := os.ReadFile("./fixtures/expected_result.csv")
			expectedResultLines := strings.Split(string(expectedResultCsv), "\n")

			Expect(formatted).To(Equal(expectedResultLines))
		})
	})

	Context("In the sad path", func() {
		It("returns an error if the header string is unexpected", func() {
			badHeaderCsv, _ := os.ReadFile("./fixtures/bad_header_n26.csv")
			badHeaderLines := strings.Split(string(badHeaderCsv), "\n")

      _, err := FormatN26(badHeaderLines)

			Expect(err.Error()).To(Equal("Unexpected header for n26 format"))
		})
	})
})
