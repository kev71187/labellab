package models_test

import (
	"app/fixtures"
	. "app/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	// "gopkg.in/jarcoal/httpmock.v1"
)

var _ = Describe("Document test", func() {
	Context("Helpers", func() {

		It("Should return a dataset", func() {
			var txt = []byte(fixtures.Dataset())
			var dataset = JsonToDataset(txt)
			Expect(dataset.Id).To(Equal(uint64(1)))
		})

		// It("Should return a dataset", func() {
		// 	httpmock.Activate()
		// 	defer httpmock.DeactivateAndReset()
		// 	httpmock.RegisterResponder("GET", "http://nginx/api/documents/1",
		// 		httpmock.NewStringResponder(200, fixtures.Document()))
		// 	Expect(GetDocument(1).Id).To(Equal(uint64(1)))
		// })
	})
})
