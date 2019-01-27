package models_test

import (
	"app/config"
	"app/fixtures"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	// "github.com/onsi/gomega/gexec"
	. "app/models"
	"gopkg.in/jarcoal/httpmock.v1"
)

var _ = Describe("Currentuser test", func() {
	Context("Helpers", func() {
		It("exe should fail", func() {
			Expect(IsOurUrl("/test/img.exe")).To(Equal(false))
		})

		It("Should return a current user", func() {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			httpmock.RegisterResponder("POST", config.BaseUrl+"users/login",
				httpmock.NewStringResponder(200, fixtures.CurrentUser()))
			var user = CurrentUserLogin("username", "password")
			Expect(user.Id).To(Equal(uint64(1)))
		})

		It("Should encrypt a string", func() {
			originalStr := "test encrypt"
			str, _ := Encrypt([]byte(originalStr))
			newStr, _ := Decrypt(str)

			Expect(string(newStr)).To(Equal(originalStr))
		})
		// 	httpmock.Activate()
		// 	defer httpmock.DeactivateAndReset()
		// 	httpmock.RegisterResponder("POST", "http://nginx/api/users/login",
		// 		httpmock.NewStringResponder(403, `{"error": "Unauthorized"}`))
		// 	Eventually(CurrentUserLogin("username", "password")).Should(gexec.Exit(1))

		// })
	})
})
