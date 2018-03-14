package match_name_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMatchName(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Match Name Suite")
}
