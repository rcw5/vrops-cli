package clients_test

import (
	"fmt"
	"net/http"

	"github.com/rcw5/vrops-cli/fakes"
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
		client = clients.NewVROpsClient(server.URL(), "hello", "world", false)
	})

	AfterEach(func() {
		server.Close()
	})

	Context("#CreateResource", func() {
		It("POSTs the resource definition to the endpoint", func() {
			resource := fakes.FakeResources[0]
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", fmt.Sprintf("/suite-api/api/resources/adapterkinds/%s", resource.ResourceKey.AdapterKindKey)),
					ghttp.VerifyBasicAuth("hello", "world"),
					ghttp.VerifyJSONRepresenting(resource),
				),
			)

			client.CreateResource(resource)
			Expect(server.ReceivedRequests()).To(HaveLen(1))
		})
	})

	Context("#CreateStats", func() {
		It("POSTs numeric stats to the endpoint", func() {
			statsForvRops := struct {
				Stats []models.Stat `json:"stat-content"`
			}{
				fakes.FakeStats,
			}
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/suite-api/api/resources/a-resource/stats"),
					ghttp.VerifyBasicAuth("hello", "world"),
					ghttp.VerifyJSONRepresenting(statsForvRops),
				),
			)

			client.CreateStats("a-resource", fakes.FakeStats)
			Expect(server.ReceivedRequests()).To(HaveLen(1))
		})
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

	Context("#ResourcesForAdapterKind", func() {
		var returnedResources models.Resources
		var returnedPageInfo models.PageInfo
		var statusCode int

		BeforeEach(func() {
			returnedResources = fakes.FakeResources
			returnedPageInfo = models.PageInfo{
				TotalCount: 1,
				PageSize:   1000,
			}
			data := struct {
				ResourceList *models.Resources `json:"resourceList"`
				PageInfo     *models.PageInfo  `json:"PageInfo"`
			}{
				ResourceList: &returnedResources,
				PageInfo:     &returnedPageInfo,
			}

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/suite-api/api/adapterkinds/my-adapterkind/resources"),
					ghttp.VerifyBasicAuth("hello", "world"),
					ghttp.RespondWithJSONEncodedPtr(&statusCode, &data),
				),
			)
		})
		Context("When the request succeeds", func() {
			BeforeEach(func() {
				statusCode = http.StatusOK
			})
			It("Returns a list of resources", func() {
				resources, err := client.ResourcesForAdapterKind("my-adapterkind")
				Expect(err).NotTo(HaveOccurred())
				Expect(resources).To(Equal(returnedResources))
			})
		})
		Context("When the adapterkind does not exist", func() {
			BeforeEach(func() {
				statusCode = http.StatusOK
				returnedResources = models.Resources{}
			})
			It("Returns no resources", func() {
				resources, err := client.ResourcesForAdapterKind("my-adapterkind")
				Expect(err).NotTo(HaveOccurred())
				Expect(resources).To(BeEmpty())
				Expect(resources).To(Equal(returnedResources))
			})
		})
		Context("When more than 1 page of information is returned", func() {
			BeforeEach(func() {
				returnedPageInfo = models.PageInfo{
					TotalCount: 5,
				}
			})
			It("Returns an error", func() {
				_, err := client.ResourcesForAdapterKind("my-adapterkind")
				Expect(err).To(MatchError("No support for result pagination yet, mate"))
			})
		})

	})

	Context("#AdapterKinds", func() {
		var returnedAdapterKinds []models.AdapterKind
		var statusCode int

		BeforeEach(func() {
			statusCode = http.StatusOK
			returnedAdapterKinds = fakes.FakeAdapterKinds
			data := struct {
				Adapters *[]models.AdapterKind `json:"adapter-kind"`
			}{
				Adapters: &returnedAdapterKinds,
			}

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/suite-api/api/adapterkinds"),
					ghttp.VerifyBasicAuth("hello", "world"),
					ghttp.RespondWithJSONEncodedPtr(&statusCode, &data),
				),
			)
		})
		Context("When the request succeeds", func() {
			It("Retrieves a list of adapter kinds", func() {
				adapters, err := client.AdapterKinds()
				Expect(err).NotTo(HaveOccurred())
				Expect(adapters).To(Equal(returnedAdapterKinds))
			})
		})

		Context("When the request fails", func() {
			BeforeEach(func() {
				statusCode = http.StatusInternalServerError
			})
			It("Returns an error", func() {
				adapters, err := client.AdapterKinds()
				Expect(err).To(HaveOccurred())
				Expect(adapters).To(BeEmpty())
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
