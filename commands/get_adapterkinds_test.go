package commands_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/topflight-technology/vrops-cli/commands"
	"github.com/topflight-technology/vrops-cli/fakes"
	"github.com/topflight-technology/vrops-cli/models"
)

var _ = Describe("GetAdapterkinds", func() {

	Context("#GetAdapterKinds", func() {
		var fakeClient fakes.FakeVRopsClient
		var fakePresenter fakes.FakePresenter

		BeforeEach(func() {
			fakeClient = fakes.FakeVRopsClient{}
			fakePresenter = fakes.FakePresenter{}
		})

		Context("When valid results are returned", func() {
			JustBeforeEach(func() {
				fakeClient.AdapterKindsReturns(fakes.FakeAdapterKinds, nil)
			})

			It("Sends the adapterkinds to the presenter", func() {
				err := commands.GetAdapterKinds(&fakeClient, &fakePresenter)
				Expect(err).NotTo(HaveOccurred())
				Expect(fakePresenter.PresentAdapterKindsCallCount()).To(Equal(1))
				Expect(fakePresenter.PresentAdapterKindsArgsForCall(0)).To(Equal(fakes.FakeAdapterKinds))
			})
		})
		Context("When the client returns an error", func() {
			It("Returns the error back to the caller", func() {
				fakeClient.AdapterKindsReturns([]models.AdapterKind{}, errors.New("An error"))
				err := commands.GetAdapterKinds(&fakeClient, &fakePresenter)
				Expect(err).To(MatchError("An error"))
				Expect(fakePresenter.PresentAdapterKindsCallCount()).To(Equal(0))

			})
		})
	})
})
