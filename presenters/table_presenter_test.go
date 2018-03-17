package presenters_test

import (
	"bytes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rcw5/vrops-cli/fakes"

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

	Context("#PresentResources", func() {
		It("Returns table encoded output", func() {
			presenter.PresentResources(fakes.FakeResources)
			Expect(buffer.String()).To(Equal(`+-------------+---------------+----------------+-----------------+-------------+--------+
|    NAME     |  IDENTIFIER   |  ADAPTERKIND   |  RESOURCEKIND   | DESCRIPTION | HEALTH |
+-------------+---------------+----------------+-----------------+-------------+--------+
| my-resource | an-identifier | my-adapterkind | my-resourcekind | Description | GREEN  |
+-------------+---------------+----------------+-----------------+-------------+--------+
`))
		})
	})

	Context("#PresentAdapterKinds", func() {
		It("Returns table encoded output", func() {
			presenter.PresentAdapterKinds(fakes.FakeAdapterKinds)
			Expect(buffer.String()).To(Equal(`+---------------+------------------+----------------------------+-----------------+
|      KEY      |       NAME       |        DESCRIPTION         | ADAPTERKINDTYPE |
+---------------+------------------+----------------------------+-----------------+
| Adapter Key   | my-adapterkind   | Nice long description here | Type            |
| Adapter Key 2 | my-adapterkind-2 | Nice long description here | Type            |
+---------------+------------------+----------------------------+-----------------+
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
