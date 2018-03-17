package commands_test

import (
	"errors"
	"time"

	"github.com/rcw5/vrops-cli/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rcw5/vrops-cli/commands"
	"github.com/rcw5/vrops-cli/fakes"
)

var _ = Describe("CreateStat", func() {

	Context("#CreateStat", func() {
		var fakeClient fakes.FakeVRopsClient
		var expectedStat models.Stats

		BeforeEach(func() {
			fakeClient = fakes.FakeVRopsClient{}
			expectedStat = models.Stats{
				models.Stat{
					StatKey:    "key1|key2",
					Data:       []float64{75},
					Timestamps: []int64{12345},
				},
			}
			fakeClient.FindResourceReturns(fakes.FakeResources[0], nil)
		})

		Context("When the client cannot find the resource", func() {
			It("Returns an error", func() {
				fakeClient.FindResourceReturns(models.Resource{}, errors.New("an error"))
				err := commands.CreateStat("my-adapterkind", "my-resource", "key1|key2", "1521282189000", 75, &fakeClient)
				Expect(err).To(MatchError("an error"))
			})
		})

		Context("When a valid date is provided in timestamp format", func() {
			JustBeforeEach(func() {
				fakeClient.CreateStatsReturns(nil)
				expectedStat[0].Timestamps[0] = 1521282189000
			})

			It("Does not return an error", func() {
				err := commands.CreateStat("my-adapterkind", "my-resource", "key1|key2", "1521282189000", 75, &fakeClient)
				Expect(err).NotTo(HaveOccurred())
				Expect(fakeClient.CreateStatsCallCount()).To(Equal(1))

				actualResource, actualStats := fakeClient.CreateStatsArgsForCall(0)
				Expect(actualResource).To(Equal("an-identifier"))
				Expect(actualStats).To(Equal(expectedStat))
			})
		})

		Context("When a valid date is provided in RFC3399 format", func() {
			JustBeforeEach(func() {
				fakeClient.CreateStatsReturns(nil)
				expectedStat[0].Timestamps[0] = 1521282189000
			})

			It("Does not return an error", func() {
				err := commands.CreateStat("my-adapterkind", "my-resource", "key1|key2", "2018-03-17T10:23:09+00:00", 75, &fakeClient)
				Expect(err).NotTo(HaveOccurred())
				Expect(fakeClient.CreateStatsCallCount()).To(Equal(1))

				actualResource, actualStats := fakeClient.CreateStatsArgsForCall(0)
				Expect(actualResource).To(Equal("an-identifier"))
				Expect(actualStats).To(Equal(expectedStat))
			})
		})

		Context("When no date is provided", func() {
			JustBeforeEach(func() {
				fakeClient.CreateStatsReturns(nil)
			})

			It("Uses the current time", func() {
				timestamp := time.Now().UnixNano() / int64(time.Millisecond)
				err := commands.CreateStat("my-adapterkind", "my-resource", "key1|key2", "", 75, &fakeClient)
				Expect(err).NotTo(HaveOccurred())
				Expect(fakeClient.CreateStatsCallCount()).To(Equal(1))

				actualResource, actualStats := fakeClient.CreateStatsArgsForCall(0)
				Expect(actualResource).To(Equal("an-identifier"))
				Expect(actualStats[0].StatKey).To(Equal("key1|key2"))
				Expect(actualStats[0].Data).To(ConsistOf([]float64{75}))
				Expect(actualStats[0].Timestamps[0]).To(Equal(timestamp))
			})
		})

		Context("When an invalid date is provided", func() {
			It("Returns an error", func() {
				err := commands.CreateStat("my-adapterkind", "my-resource", "key1|key2", "not-a-date", 75, &fakeClient)
				Expect(err).To(MatchError(`Cannot parse time: parsing time "not-a-date" as "2006-01-02T15:04:05Z07:00": cannot parse "not-a-date" as "2006"`))
			})
		})

		Context("When the client returns an error", func() {
			It("Returns the error back to the caller", func() {
				fakeClient.CreateStatsReturns(errors.New("An error"))

				err := commands.CreateStat("my-adapterkind", "my-resource", "key1|key2", "12345", 75, &fakeClient)
				Expect(err).To(MatchError("An error"))
				Expect(fakeClient.CreateStatsCallCount()).To(Equal(1))

				actualResource, actualStats := fakeClient.CreateStatsArgsForCall(0)
				Expect(actualResource).To(Equal("an-identifier"))
				Expect(actualStats).To(Equal(expectedStat))
			})
		})
	})
})
