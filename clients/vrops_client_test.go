package clients_test

import (
	"net/http"

	"github.com/rcw5/vrops-cli/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	"github.com/rcw5/vrops-cli/clients"
)

var _ = Describe("VRops Client", func() {
	var server *ghttp.Server
	var client clients.VRopsClient

	var _ = BeforeEach(func() {
		server = ghttp.NewServer()
		client = clients.NewVROpsClient(server.URL(), "hello", "world")
	})

	AfterEach(func() {
		server.Close()
	})

	Context("ResourceKinds", func() {
		var returnedAdapter models.AdapterKind
		var statusCode int

		BeforeEach(func() {
			returnedAdapter = models.AdapterKind{
				ResourceKinds: []string{"Resource1", "Resource2", "Resource3"},
			}

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/suite-api/api/adapterkinds/adapter"),
					ghttp.VerifyBasicAuth("hello", "world"),
					ghttp.RespondWithJSONEncodedPtr(&statusCode, &returnedAdapter),
				),
			)
		})
		Context("When the request succeeds", func() {
			BeforeEach(func() {
				statusCode = http.StatusOK
			})
			It("Retrieves a list of resource kinds for the adapter", func() {
				adapters, err := client.ResourceKinds("adapter")
				Expect(err).NotTo(HaveOccurred())
				Expect(adapters).To(Equal(returnedAdapter.ResourceKinds))
			})
		})

		Context("When the adapter does not exist", func() {
			BeforeEach(func() {
				statusCode = http.StatusNotFound
			})
			It("Returns an error", func() {
				adapters, err := client.ResourceKinds("adapter")
				Expect(err).To(MatchError("Request failed: 404"))
				Expect(adapters).To(BeEmpty())
			})
		})

		Context("when the request fails completely", func() {
			BeforeEach(func() {
				server.Reset()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/suite-api/api/adapterkinds/adapter"),
						func(w http.ResponseWriter, r *http.Request) {
							server.CloseClientConnections()
						},
					),
				)
			})
			It("returns an error", func() {
				_, err := client.ResourceKinds("adapter")
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Context("AdapterKinds", func() {
		var returnedAdapters models.AdapterKinds
		var statusCode int

		BeforeEach(func() {
			returnedAdapters = models.AdapterKinds{
				AdapterKind: []models.AdapterKind{
					models.AdapterKind{
						Name: "Foobar",
					},
				},
			}

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/suite-api/api/adapterkinds"),
					ghttp.VerifyBasicAuth("hello", "world"),
					ghttp.RespondWithJSONEncodedPtr(&statusCode, &returnedAdapters),
				),
			)
		})
		Context("When the request succeeds", func() {
			BeforeEach(func() {
				statusCode = http.StatusOK
			})
			It("Retrieves a list of adapter kinds", func() {
				adapters, err := client.AdapterKinds()
				Expect(err).NotTo(HaveOccurred())
				Expect(adapters).To(Equal(returnedAdapters))
			})
		})

		Context("When the request fails", func() {
			BeforeEach(func() {
				statusCode = http.StatusInternalServerError
			})
			It("Returns an error", func() {
				adapters, err := client.AdapterKinds()
				Expect(err).To(HaveOccurred())
				Expect(adapters.AdapterKind).To(BeEmpty())
			})
		})

		Context("when the request fails completely", func() {
			BeforeEach(func() {
				server.Reset()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/suite-api/api/adapterkinds"),
						func(w http.ResponseWriter, r *http.Request) {
							server.CloseClientConnections()
						},
					),
				)
			})
			It("returns an error", func() {
				_, err := client.AdapterKinds()
				Expect(err).To(HaveOccurred())
			})
		})

	})
})
