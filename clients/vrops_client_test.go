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
				Stats models.Stats `json:"stat-content"`
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

	Context("#ResourceKinds", func() {
		var returnedAdapter models.AdapterKind
		var statusCode int

		BeforeEach(func() {
			returnedAdapter = models.AdapterKind{
				ResourceKinds: []string{"Resource1", "Resource2", "Resource3"},
			}

			server.Reset()
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

		Context("When the response contains invalid JSON", func() {
			BeforeEach(func() {
				server.Reset()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/suite-api/api/adapterkinds/adapter"),
						ghttp.VerifyBasicAuth("hello", "world"),
						ghttp.RespondWith(http.StatusOK, "not json"),
					),
				)
			})
			It("Returns an error", func() {
				_, err := client.ResourceKinds("adapter")
				Expect(err).To(MatchError(ContainSubstring("Cannot parse response:")))
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
				Expect(err).To(MatchError(ContainSubstring("/suite-api/api/adapterkinds/adapter: EOF")))
			})
		})
	})

	ConfigureForAdapterAndResources := func(adapterKindsStatusCode, resourcesStatusCode *int) {
		adapterKindData := struct {
			Adapters *models.AdapterKinds `json:"adapter-kind"`
		}{
			Adapters: &fakes.FakeAdapterKinds,
		}
		returnedPageInfo := models.PageInfo{
			TotalCount: 1,
			PageSize:   1000,
		}
		resourceData := struct {
			ResourceList *models.Resources `json:"resourceList"`
			PageInfo     *models.PageInfo  `json:"PageInfo"`
		}{
			ResourceList: &fakes.FakeResources,
			PageInfo:     &returnedPageInfo,
		}
		server.AppendHandlers(
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/suite-api/api/adapterkinds"),
				ghttp.VerifyBasicAuth("hello", "world"),
				ghttp.RespondWithJSONEncodedPtr(adapterKindsStatusCode, &adapterKindData),
			),
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/suite-api/api/adapterkinds/my-adapterkind/resources"),
				ghttp.VerifyBasicAuth("hello", "world"),
				ghttp.RespondWithJSONEncodedPtr(resourcesStatusCode, &resourceData),
			),
		)

	}

	ErrorsWhenAdapterKindOrResourceNotFound := func(adapterKindsStatusCode, resourcesStatusCode *int) {
		Context("When the adapterkind cannot be found", func() {
			It("Returns an error", func() {
				_, err := client.FindResource("invalid-adapterkind", "my-resource")
				Expect(err).To(MatchError("Cannot find adapterkind: invalid-adapterkind"))
			})
		})
		Context("When the adapterkind request returns a 400", func() {
			BeforeEach(func() {
				*adapterKindsStatusCode = http.StatusBadRequest
			})
			It("Returns an error", func() {
				_, err := client.FindResource("invalid-adapterkind", "my-resource")
				Expect(err).To(MatchError("Error retrieving adapterkinds: Request failed: 400"))
			})
		})
		Context("When the adapterkind request fails altogether", func() {
			BeforeEach(func() {
				server.Reset()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/suite-api/api/adapterkinds"),
						ghttp.VerifyBasicAuth("hello", "world"),
						func(w http.ResponseWriter, r *http.Request) {
							server.CloseClientConnections()
						},
					),
				)
			})
			It("returns an error", func() {
				_, err := client.GetStatsForResource("my-adapterkind", "my-resource")
				Expect(err).To(MatchError(ContainSubstring("Error retrieving adapterkinds:")))
			})
		})
		Context("When the resource cannot be found", func() {
			It("Returns an error", func() {
				_, err := client.FindResource("my-adapterkind", "invalid-resource")
				Expect(err).To(MatchError("Cannot find resource: invalid-resource"))
			})
		})
		Context("When the resource lookup request returns a 400", func() {
			BeforeEach(func() {
				*resourcesStatusCode = http.StatusBadRequest
			})
			It("Returns an error", func() {
				_, err := client.FindResource("my-adapterkind", "invalid-resource")
				Expect(err).To(MatchError("Error retrieving resources: Request failed: 400"))
			})
		})
		Context("When the resource lookup request fails altogether", func() {
			BeforeEach(func() {
				server.Reset()
				adapterKindData := struct {
					Adapters *models.AdapterKinds `json:"adapter-kind"`
				}{
					Adapters: &fakes.FakeAdapterKinds,
				}
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/suite-api/api/adapterkinds"),
						ghttp.VerifyBasicAuth("hello", "world"),
						ghttp.RespondWithJSONEncodedPtr(adapterKindsStatusCode, &adapterKindData),
					),
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/suite-api/api/adapterkinds/my-adapterkind/resources"),
						ghttp.VerifyBasicAuth("hello", "world"),
						func(w http.ResponseWriter, r *http.Request) {
							server.CloseClientConnections()
						},
					),
				)
			})
			It("returns an error", func() {
				_, err := client.GetStatsForResource("my-adapterkind", "my-resource")
				Expect(err).To(MatchError(ContainSubstring("Error retrieving resources:")))
			})
		})
	}

	Context("#FindResource", func() {
		var adapterKindsStatusCode int
		var resourcesStatusCode int

		BeforeEach(func() {
			adapterKindsStatusCode = http.StatusOK
			resourcesStatusCode = http.StatusOK

			server.Reset()
			ConfigureForAdapterAndResources(&adapterKindsStatusCode, &resourcesStatusCode)

		})
		It("Returns the resource when found", func() {
			resource, err := client.FindResource("my-adapterkind", "my-resource")
			Expect(err).NotTo(HaveOccurred())
			Expect(resource.Identifier).To(Equal("an-identifier"))
		})

		ErrorsWhenAdapterKindOrResourceNotFound(&adapterKindsStatusCode, &resourcesStatusCode)

	})

	Context("#GetStatsForResource", func() {
		var adapterKindsStatusCode int
		var resourcesStatusCode int
		var statsStatusCode int
		var stats models.ListStatsResponse

		BeforeEach(func() {
			adapterKindsStatusCode = http.StatusOK
			resourcesStatusCode = http.StatusOK
			statsStatusCode = http.StatusOK
			server.Reset()
			ConfigureForAdapterAndResources(&adapterKindsStatusCode, &resourcesStatusCode)

			stats = models.ListStatsResponse{
				Values: []models.ListStatsResponseValues{
					models.ListStatsResponseValues{
						ResourceID: "resource-12345",
						StatList: models.ListStatsResponseValuesStatList{
							Stat: models.ListStatsResponseValuesStatListStats{
								models.ListStatsResponseValuesStatListStat{
									StatKey:      models.ListStatsResponseValuesStatListStatStatKey{Key: "Hello"},
									Timestamps:   []int64{100, 200, 300},
									Data:         []float64{1, 2, 3},
									IntervalUnit: models.IntervalUnit{Quantifier: 5},
								},
							},
						},
					},
				},
			}
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/suite-api/api/resources/an-identifier/stats"),
					ghttp.VerifyBasicAuth("hello", "world"),
					ghttp.RespondWithJSONEncodedPtr(&statsStatusCode, &stats),
				),
			)
		})

		ErrorsWhenAdapterKindOrResourceNotFound(&adapterKindsStatusCode, &resourcesStatusCode)

		It("Returns the stats for the provided resource", func() {
			stats, err := client.GetStatsForResource("my-adapterkind", "my-resource")
			Expect(err).NotTo(HaveOccurred())
			Expect(stats).To(Equal(stats))
			Expect(server.ReceivedRequests()).To(HaveLen(3))
		})

		Context("When the stats lookup fails", func() {
			BeforeEach(func() {
				statsStatusCode = http.StatusNotFound
			})
			It("returns an error", func() {
				_, err := client.GetStatsForResource("my-adapterkind", "my-resource")
				Expect(err).To(MatchError("Request failed: 404"))
			})
		})

		Context("When the stats lookup returns bad JSON", func() {
			BeforeEach(func() {
				server.Reset()
				ConfigureForAdapterAndResources(&adapterKindsStatusCode, &resourcesStatusCode)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/suite-api/api/resources/an-identifier/stats"),
						ghttp.RespondWith(http.StatusOK, "{bad_json}"),
					),
				)
			})
			It("returns an error", func() {
				_, err := client.GetStatsForResource("my-adapterkind", "my-resource")
				Expect(err).To(MatchError(ContainSubstring("Cannot parse response:")))
			})
		})

		Context("When the stats request fails altogether", func() {
			BeforeEach(func() {
				server.Reset()
				ConfigureForAdapterAndResources(&adapterKindsStatusCode, &resourcesStatusCode)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/suite-api/api/resources/an-identifier/stats"),
						func(w http.ResponseWriter, r *http.Request) {
							server.CloseClientConnections()
						},
					),
				)
			})
			It("returns an error", func() {
				_, err := client.GetStatsForResource("my-adapterkind", "my-resource")
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
		var returnedAdapterKinds models.AdapterKinds
		var statusCode int

		BeforeEach(func() {
			statusCode = http.StatusOK
			returnedAdapterKinds = fakes.FakeAdapterKinds
			data := struct {
				Adapters *models.AdapterKinds `json:"adapter-kind"`
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
