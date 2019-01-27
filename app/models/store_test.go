package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	// "github.com/onsi/gomega/gexec"
	. "app/models"
)

var _ = Describe("Store test", func() {
	Context("Helpers", func() {
		It("Should encrypt a string", func() {
			originalStr := "test encrypt"
			str, _ := Encrypt([]byte(originalStr))
			newStr, _ := Decrypt(str)

			Expect(string(newStr)).To(Equal(originalStr))
		})
	})
})
