package models_test

import (
	"app/config"
	"app/fixtures"
	"app/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/jarcoal/httpmock.v1"
)

var _ = Describe("Dataset test", func() {
	Context("Helpers", func() {
		It("exe should fail", func() {
			Expect(models.IsOurUrl("/test/img.exe")).To(Equal(false))
		})

		It("jpg should pass", func() {
			Expect(models.FileTypeMatch("/test/img.jpg")).To(Equal(true))
		})

		It("Should hash our file with md5", func() {
			Expect(models.GetMD5Hash([]byte("Here is a string...."))).To(Equal("46d55ab4c610163c6e25a3109d0a0506"))
		})

		It("Should hash our file with md5", func() {
			Expect(models.GetMD5Hash([]byte("Here is a string..."))).To(Equal("5714a31a9d70f3bc5d5cc1a62b8cd606"))
		})

		It("Should parse a int", func() {
			Expect(models.IdToString(123)).To(Equal("123"))
		})

		It("Should return a dataset", func() {
			var txt = []byte(fixtures.Dataset())
			var dataset = models.JsonToDataset(txt)
			Expect(dataset.Id).To(Equal(uint64(1)))
		})

		It("Should return a dataset", func() {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			httpmock.RegisterResponder("GET", config.BaseUrl+"datasets/1",
				httpmock.NewStringResponder(200, fixtures.Dataset()))
			Expect(models.GetDataset(1).Id).To(Equal(uint64(1)))
		})
	})
})
