package commands_test

import (
	"encoding/json"
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rcw5/vrops-cli/commands"
	"github.com/rcw5/vrops-cli/fakes"
)

var _ = Describe("CreateStats", func() {

	Context("#CreateStats", func() {
		var fakeClient fakes.FakeVRopsClient
		var statsJson string

		BeforeEach(func() {
			fakeClient = fakes.FakeVRopsClient{}
			statsByteArr, err := json.Marshal(fakes.FakeStats)
			Expect(err).NotTo(HaveOccurred())
			statsJson = string(statsByteArr)
		})

		Context("When the command succeeds", func() {
			JustBeforeEach(func() {
				fakeClient.CreateStatsReturns(nil)
			})

			It("Does not return an error", func() {
				err := commands.CreateStats("a-resource", statsJson, &fakeClient)
				Expect(err).NotTo(HaveOccurred())
				Expect(fakeClient.CreateStatsCallCount()).To(Equal(1))
				actualResource, actualStats := fakeClient.CreateStatsArgsForCall(0)
				Expect(actualResource).To(Equal("a-resource"))
				Expect(actualStats).To(Equal(fakes.FakeStats))
			})
		})
		Context("When the provided JSON does not parse correctly", func() {
			It("Returns an error", func() {
				err := commands.CreateStats("a-resource", "{}", &fakeClient)
				Expect(err).To(MatchError("json: cannot unmarshal object into Go value of type []models.Stat"))
			})
		})
		Context("When the client returns an error", func() {
			It("Returns the error back to the caller", func() {
				fakeClient.CreateStatsReturns(errors.New("An error"))

				err := commands.CreateStats("a-resource", statsJson, &fakeClient)
				Expect(err).To(MatchError("An error"))
				Expect(fakeClient.CreateStatsCallCount()).To(Equal(1))
				actualResource, actualStats := fakeClient.CreateStatsArgsForCall(0)
				Expect(actualResource).To(Equal("a-resource"))
				Expect(actualStats).To(Equal(fakes.FakeStats))
			})
		})
	})
})
