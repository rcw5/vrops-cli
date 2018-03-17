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

	Context("#AdapterKinds.FindAdapterKind", func() {
		var adapterkinds AdapterKinds
		JustBeforeEach(func() {
			adapterkinds = fakes.FakeAdapterKinds
		})
		It("Returns the resource if it exists", func() {
			adapterkind, err := adapterkinds.FindAdapterKind("my-adapterkind")
			Expect(err).NotTo(HaveOccurred())
			Expect(adapterkind.Name).To(Equal("my-adapterkind"))
		})
		It("Returns an error if it does not exist", func() {
			_, err := adapterkinds.FindAdapterKind("another-adapterkind")
			Expect(err).To(MatchError("Cannot find adapterkind: another-adapterkind"))
		})
	})
})
