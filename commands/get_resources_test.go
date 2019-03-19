package commands_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/topflight-technology/vrops-cli/commands"
	"github.com/topflight-technology/vrops-cli/fakes"
)

var _ = Describe("GetResources", func() {

	Context("#GetResources", func() {
		var fakeClient fakes.FakeVRopsClient
		var fakePresenter fakes.FakePresenter

		BeforeEach(func() {
			fakeClient = fakes.FakeVRopsClient{}
			fakePresenter = fakes.FakePresenter{}
		})

		Context("When valid results are returned", func() {
			JustBeforeEach(func() {
				fakeClient.ResourcesForAdapterKindReturns(fakes.FakeResources, nil)
			})

			It("Sends the resources to the presenter", func() {
				err := commands.GetResources("an-adapter", &fakeClient, &fakePresenter)
				Expect(err).NotTo(HaveOccurred())
				Expect(fakeClient.ResourcesForAdapterKindCallCount()).To(Equal(1))
				Expect(fakeClient.ResourcesForAdapterKindArgsForCall(0)).To(Equal("an-adapter"))
				Expect(fakePresenter.PresentResourcesCallCount()).To(Equal(1))
				Expect(fakePresenter.PresentResourcesArgsForCall(0)).To(Equal(fakes.FakeResources))
			})
		})
		Context("When the client returns an error", func() {
			It("Returns the error back to the caller", func() {
				fakeClient.ResourcesForAdapterKindReturns(nil, errors.New("An error"))
				err := commands.GetResources("an-adapter", &fakeClient, &fakePresenter)
				Expect(err).To(MatchError("An error"))
				Expect(fakePresenter.PresentResourcesCallCount()).To(Equal(0))

			})
		})
	})
})
