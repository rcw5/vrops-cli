package commands_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/topflight-technology/vrops-cli/commands"
	"github.com/topflight-technology/vrops-cli/fakes"
)

var _ = Describe("GetResourceKinds", func() {

	Context("#GetResourceKiunds", func() {
		var fakeClient fakes.FakeVRopsClient
		var fakePresenter fakes.FakePresenter
		var resourceKinds []string

		BeforeEach(func() {
			fakeClient = fakes.FakeVRopsClient{}
			fakePresenter = fakes.FakePresenter{}
		})

		Context("When valid results are returned", func() {
			JustBeforeEach(func() {
				resourceKinds = []string{"Resource1", "Resource2"}
				fakeClient.ResourceKindsReturns(resourceKinds, nil)
			})

			It("Sends the resource kinds to the presenter", func() {
				err := commands.GetResourceKinds("an-adapter", &fakeClient, &fakePresenter)
				Expect(err).NotTo(HaveOccurred())
				Expect(fakeClient.ResourceKindsCallCount()).To(Equal(1))
				Expect(fakeClient.ResourceKindsArgsForCall(0)).To(Equal("an-adapter"))
				Expect(fakePresenter.PresentResourceKindsCallCount()).To(Equal(1))
				Expect(fakePresenter.PresentResourceKindsArgsForCall(0)).To(Equal(resourceKinds))
			})
		})
		Context("When the client returns an error", func() {
			It("Returns the error back to the caller", func() {
				fakeClient.ResourceKindsReturns(nil, errors.New("An error"))
				err := commands.GetResourceKinds("an-adapter", &fakeClient, &fakePresenter)
				Expect(err).To(MatchError("An error"))
				Expect(fakePresenter.PresentAdapterKindsCallCount()).To(Equal(0))

			})
		})
	})
})
