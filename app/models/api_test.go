package models

import (
	"app/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Api test", func() {
	Context("should pass", func() {
		It("not ours should fail", func() {
			var ours = IsOurUrl("https://test.com")
			Expect(ours).To(Equal(false))
		})

		It("ours should pass", func() {
			Expect(IsOurUrl(config.BaseUrl)).To(Equal(true))
		})
	})
})
