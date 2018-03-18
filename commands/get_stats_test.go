package commands_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rcw5/vrops-cli/commands"
	"github.com/rcw5/vrops-cli/fakes"
)

var _ = Describe("Get Stats", func() {

	Context("#GetStats", func() {
		var fakeClient fakes.FakeVRopsClient
		var fakePresenter fakes.FakePresenter

		BeforeEach(func() {
			fakeClient = fakes.FakeVRopsClient{}
			fakePresenter = fakes.FakePresenter{}
		})

		Context("When valid results are returned", func() {
			JustBeforeEach(func() {
				fakeClient.GetStatsForResourceReturns(fakes.FakeListStatsResponse, nil)
			})
			Context("and no stat-key is provided", func() {
				It("Sends the resources to the presenter", func() {
					err := commands.GetStats("an-adapter", "a-resource", "", &fakeClient, &fakePresenter)
					Expect(err).NotTo(HaveOccurred())
					Expect(fakeClient.GetStatsForResourceCallCount()).To(Equal(1))
					adapter, resource, statKey := fakeClient.GetStatsForResourceArgsForCall(0)
					Expect(adapter).To(Equal("an-adapter"))
					Expect(resource).To(Equal("a-resource"))
					Expect(statKey).To(BeEmpty())
					Expect(fakePresenter.PresentStatsCallCount()).To(Equal(1))
					Expect(fakePresenter.PresentStatsArgsForCall(0)).To(Equal(fakes.FakeListStatsResponse))
				})
			})
			Context("and a stat-key is provided", func() {
				It("Sends the resources to the presenter", func() {
					err := commands.GetStats("an-adapter", "a-resource", "my-stat", &fakeClient, &fakePresenter)
					Expect(err).NotTo(HaveOccurred())
					Expect(fakeClient.GetStatsForResourceCallCount()).To(Equal(1))
					adapter, resource, statKey := fakeClient.GetStatsForResourceArgsForCall(0)
					Expect(adapter).To(Equal("an-adapter"))
					Expect(resource).To(Equal("a-resource"))
					Expect(statKey).To(Equal("my-stat"))
					Expect(fakePresenter.PresentStatsCallCount()).To(Equal(1))
					Expect(fakePresenter.PresentStatsArgsForCall(0)).To(Equal(fakes.FakeListStatsResponse))
				})
			})
		})
		Context("When the client returns an error", func() {
			It("Returns the error back to the caller", func() {
				fakeClient.GetStatsForResourceReturns(nil, errors.New("An error"))
				err := commands.GetStats("an-adapter", "a-resource", "", &fakeClient, &fakePresenter)
				Expect(err).To(MatchError("An error"))
				Expect(fakePresenter.PresentStatsCallCount()).To(Equal(0))

			})
		})
	})
})
