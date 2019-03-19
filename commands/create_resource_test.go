package commands_test

import (
	"encoding/json"
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/topflight-technology/vrops-cli/commands"
	"github.com/topflight-technology/vrops-cli/fakes"
)

var _ = Describe("Create Resource Command", func() {

	Context("#CreateResource", func() {
		var fakeClient fakes.FakeVRopsClient
		var definition string

		BeforeEach(func() {
			fakeClient = fakes.FakeVRopsClient{}
			statsByteArr, err := json.Marshal(fakes.FakeResources[0])
			Expect(err).NotTo(HaveOccurred())
			definition = string(statsByteArr)
		})

		Context("When the command succeeds", func() {
			JustBeforeEach(func() {
				fakeClient.CreateResourceReturns(nil)
			})

			It("Does not return an error", func() {
				err := commands.CreateResource(definition, &fakeClient)
				Expect(err).NotTo(HaveOccurred())
				Expect(fakeClient.CreateResourceCallCount()).To(Equal(1))
				actualResource := fakeClient.CreateResourceArgsForCall(0)
				Expect(actualResource).To(Equal(fakes.FakeResources[0]))
			})
		})
		Context("When the provided JSON does not parse correctly", func() {
			It("Returns an error", func() {
				err := commands.CreateResource(`{ss: "json"}`, &fakeClient)
				Expect(err).To(MatchError("invalid character 's' looking for beginning of object key string"))
			})
		})
		Context("When the client returns an error", func() {
			It("Returns the error back to the caller", func() {
				fakeClient.CreateResourceReturns(errors.New("An error"))

				err := commands.CreateResource(definition, &fakeClient)
				Expect(err).To(MatchError("An error"))
				Expect(fakeClient.CreateResourceCallCount()).To(Equal(1))
				actualResource := fakeClient.CreateResourceArgsForCall(0)
				Expect(actualResource).To(Equal(fakes.FakeResources[0]))
			})
		})
	})
})
