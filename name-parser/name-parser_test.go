package name_parser_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestNameParser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Name Parser Suite")
}
