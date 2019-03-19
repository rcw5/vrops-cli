package presenters_test

import (
	"bytes"

	"github.com/topflight-technology/vrops-cli/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/topflight-technology/vrops-cli/fakes"

	. "github.com/topflight-technology/vrops-cli/presenters"
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

		Context("#PresentStats", func() {
			Context("When no stats are returned", func() {
				It("does not display a table", func() {
					presenter.PresentStats(models.ListStatsResponseValuesStatListStats{})
					Expect(buffer.String()).To(Equal("No stats found\n"))
				})
			})
			It("Returns table encoded output", func() {
				presenter.PresentStats(fakes.FakeListStatsResponse)
				Expect(buffer.String()).To(Equal(`+------------------+-----------------------------------+-------+
|       NAME       |               TIME                | VALUE |
+------------------+-----------------------------------+-------+
| stat|key         | 1970-01-01 01:00:00.001 +0100 BST |     1 |
| stat|key         | 1970-01-01 01:00:00.002 +0100 BST |     2 |
| stat|key         | 1970-01-01 01:00:00.003 +0100 BST |     3 |
| stat|key         | 1970-01-01 01:00:00.004 +0100 BST |     4 |
| stat|key         | 1970-01-01 01:00:00.005 +0100 BST |     5 |
| another-stat|key | 1970-01-01 01:00:00.001 +0100 BST |     1 |
| another-stat|key | 1970-01-01 01:00:00.002 +0100 BST |     2 |
| another-stat|key | 1970-01-01 01:00:00.003 +0100 BST |     3 |
| another-stat|key | 1970-01-01 01:00:00.004 +0100 BST |     4 |
| another-stat|key | 1970-01-01 01:00:00.005 +0100 BST |     5 |
+------------------+-----------------------------------+-------+
`))
			})
		})
	})
})
