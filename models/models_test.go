package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rcw5/vrops-cli/fakes"
	. "github.com/rcw5/vrops-cli/models"
)

var _ = Describe("Models", func() {

	Context("#Resources.FindResource", func() {
		var resources Resources
		JustBeforeEach(func() {
			resources = fakes.FakeResources
		})
		It("Returns the resource if it exists", func() {
			resource, err := resources.FindResource("my-resource")
			Expect(err).NotTo(HaveOccurred())
			Expect(resource.ResourceKey.Name).To(Equal("my-resource"))
		})
		It("Returns an error if it does not exist", func() {
			_, err := resources.FindResource("another-resource")
			Expect(err).To(MatchError("Cannot find resource: another-resource"))
		})
	})
})
