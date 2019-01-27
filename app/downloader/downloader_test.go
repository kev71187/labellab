package downloader_test

import (
	// "app/fixtures"
	. "app/downloader"
	. "app/fixtures"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	// "gopkg.in/jarcoal/httpmock.v1"
)

var _ = Describe("Download", func() {
	Context("Download test", func() {
		It("Should download images", func() {
			Expect(Download("data/", DefaultDownloadConfig())).To(Equal(1))
		})
	})
})
