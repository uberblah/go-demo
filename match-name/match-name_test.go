package match_name_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/golang/mock/gomock"
	"github.com/uberblah/go-demo/match-name"
	"github.com/uberblah/go-demo/match-name/mock/name-parser"
	"github.com/uberblah/go-demo/name-parser"
)

var _ = Describe("Match Name package", func() {

	Context("The MatchName function", func() {

		var ctrl *gomock.Controller
		var np *mock_np.MockNameParser
		var oldReg map[string]name_parser.NameParser

		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT(0))
			np = mock_np.NewMockNameParser(ctrl)

			oldReg = match_name.Registry
		})

		AfterEach(func() {
			match_name.Registry = oldReg

			ctrl.Finish()
		})

		It("Should return nil and empty string when it can't identify a name", func() {
			match_name.Registry = nil

			t, parts := match_name.MatchName("who cares what")

			Expect(t).To(Equal(""))
			Expect(parts).To(BeNil())
		})

		It("Should return nil and empty string when it can't identify a name", func() {
			parserName := "the_one_parser_to_rule_them_all"
			match_name.Registry = map[string]name_parser.NameParser{
				parserName: np,
			}

			nameToParse := "who cares what"
			correctParts := map[string]string{"hi": "there", "world": "yay"}

			gomock.InOrder(
				np.EXPECT().ParseFrom(nameToParse).Return(nil),
				np.EXPECT().GetParts().Return(correctParts),
			)

			t, parts := match_name.MatchName(nameToParse)

			Expect(t).To(Equal(parserName))
			Expect(parts).To(Equal(correctParts))
		})

	})

})
