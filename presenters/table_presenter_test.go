package presenters_test

import (
	"bytes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rcw5/vrops-cli/models"

	. "github.com/rcw5/vrops-cli/presenters"
)

var _ = Describe("TablePresenter", func() {
	var buffer bytes.Buffer
	var presenter TablePresenter

	BeforeEach(func() {
		buffer = bytes.Buffer{}
		presenter = TablePresenter{
			Buffer: &buffer,
		}
	})

	Context("#PresentAdapterKinds", func() {
		It("Returns table encoded output", func() {
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
			Expect(buffer.String()).To(Equal(`+-----+------+-------------+-----------------+
| KEY | NAME | DESCRIPTION | ADAPTERKINDTYPE |
+-----+------+-------------+-----------------+
| key | name | description | type            |
+-----+------+-------------+-----------------+
`))
		})
	})

	Context("#PresentResourceKinds", func() {
		It("Returns table encoded output", func() {
			resourceKinds := []string{"res1", "res2", "res3"}

			presenter.PresentResourceKinds(resourceKinds)
			Expect(buffer.String()).To(Equal(`+------+
| NAME |
+------+
| res1 |
| res2 |
| res3 |
+------+
`))

		})
	})
})
