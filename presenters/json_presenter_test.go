package presenters_test

import (
	"bytes"
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rcw5/vrops-cli/fakes"

	. "github.com/rcw5/vrops-cli/presenters"
)

var _ = Describe("JsonPresenter", func() {
	var buffer bytes.Buffer
	var presenter JSONPresenter

	BeforeEach(func() {
		buffer = bytes.Buffer{}
		presenter = JSONPresenter{
			Buffer: &buffer,
		}
	})

	Context("#PresentResources", func() {
		It("Returns JSON encoded output", func() {
			presenter.PresentResources(fakes.FakeResources)
			jsonData, err := json.Marshal(fakes.FakeResources)
			Expect(err).NotTo(HaveOccurred())
			Expect(buffer.String()).To(Equal(string(jsonData)))
		})
	})

	Context("#PresentAdapterKinds", func() {
		It("Returns JSON encoded output", func() {
			presenter.PresentAdapterKinds(fakes.FakeAdapterKinds)
			jsonData, err := json.Marshal(fakes.FakeAdapterKinds)
			Expect(err).NotTo(HaveOccurred())
			Expect(buffer.String()).To(Equal(string(jsonData)))
		})
	})

	Context("#PresentStats", func() {
		It("Returns JSON encoded output", func() {
			presenter.PresentStats(fakes.FakeListStatsResponse)
			jsonData, err := json.Marshal(fakes.FakeListStatsResponse)
			Expect(err).NotTo(HaveOccurred())
			Expect(buffer.String()).To(Equal(string(jsonData)))
		})
	})

	Context("#PresentResourceKinds", func() {
		It("Returns JSON encoded output", func() {
			resourceKinds := []string{"res1", "res2", "res3"}

			presenter.PresentResourceKinds(resourceKinds)
			jsonData, err := json.Marshal(resourceKinds)
			Expect(err).NotTo(HaveOccurred())
			Expect(buffer.String()).To(Equal(string(jsonData)))

		})
	})
})
