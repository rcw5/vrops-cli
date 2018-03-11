package presenters_test

import (
	"bytes"
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rcw5/vrops-cli/models"

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

	Context("#PresentAdapterKinds", func() {
		It("Returns JSON encoded output", func() {
			adapterKinds := models.AdapterKinds{
				AdapterKind: []models.AdapterKind{
					models.AdapterKind{
						AdapterKindType: "type",
						DescribeVersion: 1,
						Description:     "description",
						Key:             "key",
						Name:            "name",
						ResourceKinds:   []string{"res1", "res2"},
					},
				},
			}

			presenter.PresentAdapterKinds(adapterKinds)
			jsonData, err := json.Marshal(adapterKinds)
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
