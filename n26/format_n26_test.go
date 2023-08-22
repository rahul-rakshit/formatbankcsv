package n26

import (
	"os"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("formatN26", func() {
	It("says hello", func() {
		sampleN26Csv, _ := os.ReadFile("./fixtures/sample_n26.csv")
		sampleN26Lines := strings.Split(string(sampleN26Csv), "\n")

    formatted := FormatN26(sampleN26Lines)

		expectedResultCsv, _ := os.ReadFile("./fixtures/expected_result.csv")
		expectedResultLines := strings.Split(string(expectedResultCsv), "\n")

		Expect(formatted).To(Equal(expectedResultLines))
	})
})
