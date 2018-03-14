package name_parser_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/uberblah/go-demo/name-parser"
)

var _ = Describe("The Email Name Parser", func() {

	Context("Email::ParseFrom", func() {

		email_parser := &name_parser.Email{}

		It("Should reject things that are not emails", func() {
			err := email_parser.ParseFrom("^*&^*&^(^(^()))")
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(ContainSubstring("Could not parse"))
		})

		It("Should properly parse things that are emails", func() {
			err := email_parser.ParseFrom("uberblah@uberblah.com")
			Expect(err).To(BeNil())
			Expect(map[string]string(*email_parser)).To(Equal(map[string]string{
				"username":    "uberblah",
				"domain-name": "uberblah.com",
			}))
		})

	})

})
