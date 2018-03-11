package commands_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rcw5/vrops-cli/commands"
	"github.com/rcw5/vrops-cli/fakes"
	"github.com/rcw5/vrops-cli/models"
)

var _ = Describe("GetAdapterkinds", func() {

	Context("#GetAdapterKinds", func() {
		var fakeClient fakes.FakeVRopsClient
		var fakePresenter fakes.FakePresenter
		var adapterKinds models.AdapterKinds

		BeforeEach(func() {
			fakeClient = fakes.FakeVRopsClient{}
			fakePresenter = fakes.FakePresenter{}
		})

		Context("When valid results are returned", func() {
			JustBeforeEach(func() {
				adapterKinds = models.AdapterKinds{
					AdapterKind: []models.AdapterKind{
						models.AdapterKind{
							Key:             "Adapter Key",
							Description:     "Nice long description here",
							AdapterKindType: "Type",
							Name:            "An Adapter",
						},
					},
				}
				fakeClient.AdapterKindsReturns(adapterKinds, nil)
			})

			It("Sends the adapterkinds to the presenter", func() {
				err := commands.GetAdapterKinds(&fakeClient, &fakePresenter)
				Expect(err).NotTo(HaveOccurred())
				Expect(fakePresenter.PresentAdapterKindsCallCount()).To(Equal(1))
				Expect(fakePresenter.PresentAdapterKindsArgsForCall(0)).To(Equal(adapterKinds))
			})
		})
		Context("When the client returns an error", func() {
			It("Returns the error back to the caller", func() {
				fakeClient.AdapterKindsReturns(models.AdapterKinds{}, errors.New("An error"))
				err := commands.GetAdapterKinds(&fakeClient, &fakePresenter)
				Expect(err).To(MatchError("An error"))
				Expect(fakePresenter.PresentAdapterKindsCallCount()).To(Equal(0))

			})
		})
	})
})
