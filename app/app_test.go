package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

)

var _ = Describe("App", func() {
	Context("should pass", func() {
		It("should pass", func() {
			Expect(true).To(Equal(true))
		})
	})
})
