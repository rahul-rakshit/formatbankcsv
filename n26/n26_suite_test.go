package n26

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestN26(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "N26 Suite")
}
