package models_test

import (
	"app/config"
	"app/fixtures"
	. "app/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/jarcoal/httpmock.v1"
)

var _ = Describe("Batch", func() {
	Context("Batch", func() {
		It("Should create a batch and respond with a parsed batch", func() {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			httpmock.RegisterResponder("POST", config.BaseUrl+"batches?dataset_id=1",
				httpmock.NewStringResponder(200, fixtures.Batch()))
			Expect(CreateBatch(1).Id).To(Equal(uint64(1)))
		})
	})
})
