// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/topflight-technology/vrops-cli/clients"
	"github.com/topflight-technology/vrops-cli/models"
)

type FakeVRopsClient struct {
	AdapterKindsStub        func() (models.AdapterKinds, error)
	adapterKindsMutex       sync.RWMutex
	adapterKindsArgsForCall []struct{}
	adapterKindsReturns     struct {
		result1 models.AdapterKinds
		result2 error
	}
	adapterKindsReturnsOnCall map[int]struct {
		result1 models.AdapterKinds
		result2 error
	}
	ResourceKindsStub        func(string) ([]string, error)
	resourceKindsMutex       sync.RWMutex
	resourceKindsArgsForCall []struct {
		arg1 string
	}
	resourceKindsReturns struct {
		result1 []string
		result2 error
	}
	resourceKindsReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	ResourcesForAdapterKindStub        func(string) (models.Resources, error)
	resourcesForAdapterKindMutex       sync.RWMutex
	resourcesForAdapterKindArgsForCall []struct {
		arg1 string
	}
	resourcesForAdapterKindReturns struct {
		result1 models.Resources
		result2 error
	}
	resourcesForAdapterKindReturnsOnCall map[int]struct {
		result1 models.Resources
		result2 error
	}
	FindResourceStub        func(string, string) (models.Resource, error)
	findResourceMutex       sync.RWMutex
	findResourceArgsForCall []struct {
		arg1 string
		arg2 string
	}
	findResourceReturns struct {
		result1 models.Resource
		result2 error
	}
	findResourceReturnsOnCall map[int]struct {
		result1 models.Resource
		result2 error
	}
	CreateStatsStub        func(string, models.Stats) error
	createStatsMutex       sync.RWMutex
	createStatsArgsForCall []struct {
		arg1 string
		arg2 models.Stats
	}
	createStatsReturns struct {
		result1 error
	}
	createStatsReturnsOnCall map[int]struct {
		result1 error
	}
	GetStatsForResourceStub        func(string, string, string) (models.ListStatsResponseValuesStatListStats, error)
	getStatsForResourceMutex       sync.RWMutex
	getStatsForResourceArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	getStatsForResourceReturns struct {
		result1 models.ListStatsResponseValuesStatListStats
		result2 error
	}
	getStatsForResourceReturnsOnCall map[int]struct {
		result1 models.ListStatsResponseValuesStatListStats
		result2 error
	}
	CreateResourceStub        func(models.Resource) error
	createResourceMutex       sync.RWMutex
	createResourceArgsForCall []struct {
		arg1 models.Resource
	}
	createResourceReturns struct {
		result1 error
	}
	createResourceReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVRopsClient) AdapterKinds() (models.AdapterKinds, error) {
	fake.adapterKindsMutex.Lock()
	ret, specificReturn := fake.adapterKindsReturnsOnCall[len(fake.adapterKindsArgsForCall)]
	fake.adapterKindsArgsForCall = append(fake.adapterKindsArgsForCall, struct{}{})
	fake.recordInvocation("AdapterKinds", []interface{}{})
	fake.adapterKindsMutex.Unlock()
	if fake.AdapterKindsStub != nil {
		return fake.AdapterKindsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.adapterKindsReturns.result1, fake.adapterKindsReturns.result2
}

func (fake *FakeVRopsClient) AdapterKindsCallCount() int {
	fake.adapterKindsMutex.RLock()
	defer fake.adapterKindsMutex.RUnlock()
	return len(fake.adapterKindsArgsForCall)
}

func (fake *FakeVRopsClient) AdapterKindsReturns(result1 models.AdapterKinds, result2 error) {
	fake.AdapterKindsStub = nil
	fake.adapterKindsReturns = struct {
		result1 models.AdapterKinds
		result2 error
	}{result1, result2}
}

func (fake *FakeVRopsClient) AdapterKindsReturnsOnCall(i int, result1 models.AdapterKinds, result2 error) {
	fake.AdapterKindsStub = nil
	if fake.adapterKindsReturnsOnCall == nil {
		fake.adapterKindsReturnsOnCall = make(map[int]struct {
			result1 models.AdapterKinds
			result2 error
		})
	}
	fake.adapterKindsReturnsOnCall[i] = struct {
		result1 models.AdapterKinds
		result2 error
	}{result1, result2}
}

func (fake *FakeVRopsClient) ResourceKinds(arg1 string) ([]string, error) {
	fake.resourceKindsMutex.Lock()
	ret, specificReturn := fake.resourceKindsReturnsOnCall[len(fake.resourceKindsArgsForCall)]
	fake.resourceKindsArgsForCall = append(fake.resourceKindsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ResourceKinds", []interface{}{arg1})
	fake.resourceKindsMutex.Unlock()
	if fake.ResourceKindsStub != nil {
		return fake.ResourceKindsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.resourceKindsReturns.result1, fake.resourceKindsReturns.result2
}

func (fake *FakeVRopsClient) ResourceKindsCallCount() int {
	fake.resourceKindsMutex.RLock()
	defer fake.resourceKindsMutex.RUnlock()
	return len(fake.resourceKindsArgsForCall)
}

func (fake *FakeVRopsClient) ResourceKindsArgsForCall(i int) string {
	fake.resourceKindsMutex.RLock()
	defer fake.resourceKindsMutex.RUnlock()
	return fake.resourceKindsArgsForCall[i].arg1
}

func (fake *FakeVRopsClient) ResourceKindsReturns(result1 []string, result2 error) {
	fake.ResourceKindsStub = nil
	fake.resourceKindsReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeVRopsClient) ResourceKindsReturnsOnCall(i int, result1 []string, result2 error) {
	fake.ResourceKindsStub = nil
	if fake.resourceKindsReturnsOnCall == nil {
		fake.resourceKindsReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.resourceKindsReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeVRopsClient) ResourcesForAdapterKind(arg1 string) (models.Resources, error) {
	fake.resourcesForAdapterKindMutex.Lock()
	ret, specificReturn := fake.resourcesForAdapterKindReturnsOnCall[len(fake.resourcesForAdapterKindArgsForCall)]
	fake.resourcesForAdapterKindArgsForCall = append(fake.resourcesForAdapterKindArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ResourcesForAdapterKind", []interface{}{arg1})
	fake.resourcesForAdapterKindMutex.Unlock()
	if fake.ResourcesForAdapterKindStub != nil {
		return fake.ResourcesForAdapterKindStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.resourcesForAdapterKindReturns.result1, fake.resourcesForAdapterKindReturns.result2
}

func (fake *FakeVRopsClient) ResourcesForAdapterKindCallCount() int {
	fake.resourcesForAdapterKindMutex.RLock()
	defer fake.resourcesForAdapterKindMutex.RUnlock()
	return len(fake.resourcesForAdapterKindArgsForCall)
}

func (fake *FakeVRopsClient) ResourcesForAdapterKindArgsForCall(i int) string {
	fake.resourcesForAdapterKindMutex.RLock()
	defer fake.resourcesForAdapterKindMutex.RUnlock()
	return fake.resourcesForAdapterKindArgsForCall[i].arg1
}

func (fake *FakeVRopsClient) ResourcesForAdapterKindReturns(result1 models.Resources, result2 error) {
	fake.ResourcesForAdapterKindStub = nil
	fake.resourcesForAdapterKindReturns = struct {
		result1 models.Resources
		result2 error
	}{result1, result2}
}

func (fake *FakeVRopsClient) ResourcesForAdapterKindReturnsOnCall(i int, result1 models.Resources, result2 error) {
	fake.ResourcesForAdapterKindStub = nil
	if fake.resourcesForAdapterKindReturnsOnCall == nil {
		fake.resourcesForAdapterKindReturnsOnCall = make(map[int]struct {
			result1 models.Resources
			result2 error
		})
	}
	fake.resourcesForAdapterKindReturnsOnCall[i] = struct {
		result1 models.Resources
		result2 error
	}{result1, result2}
}

func (fake *FakeVRopsClient) FindResource(arg1 string, arg2 string) (models.Resource, error) {
	fake.findResourceMutex.Lock()
	ret, specificReturn := fake.findResourceReturnsOnCall[len(fake.findResourceArgsForCall)]
	fake.findResourceArgsForCall = append(fake.findResourceArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("FindResource", []interface{}{arg1, arg2})
	fake.findResourceMutex.Unlock()
	if fake.FindResourceStub != nil {
		return fake.FindResourceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findResourceReturns.result1, fake.findResourceReturns.result2
}

func (fake *FakeVRopsClient) FindResourceCallCount() int {
	fake.findResourceMutex.RLock()
	defer fake.findResourceMutex.RUnlock()
	return len(fake.findResourceArgsForCall)
}

func (fake *FakeVRopsClient) FindResourceArgsForCall(i int) (string, string) {
	fake.findResourceMutex.RLock()
	defer fake.findResourceMutex.RUnlock()
	return fake.findResourceArgsForCall[i].arg1, fake.findResourceArgsForCall[i].arg2
}

func (fake *FakeVRopsClient) FindResourceReturns(result1 models.Resource, result2 error) {
	fake.FindResourceStub = nil
	fake.findResourceReturns = struct {
		result1 models.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeVRopsClient) FindResourceReturnsOnCall(i int, result1 models.Resource, result2 error) {
	fake.FindResourceStub = nil
	if fake.findResourceReturnsOnCall == nil {
		fake.findResourceReturnsOnCall = make(map[int]struct {
			result1 models.Resource
			result2 error
		})
	}
	fake.findResourceReturnsOnCall[i] = struct {
		result1 models.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeVRopsClient) CreateStats(arg1 string, arg2 models.Stats) error {
	fake.createStatsMutex.Lock()
	ret, specificReturn := fake.createStatsReturnsOnCall[len(fake.createStatsArgsForCall)]
	fake.createStatsArgsForCall = append(fake.createStatsArgsForCall, struct {
		arg1 string
		arg2 models.Stats
	}{arg1, arg2})
	fake.recordInvocation("CreateStats", []interface{}{arg1, arg2})
	fake.createStatsMutex.Unlock()
	if fake.CreateStatsStub != nil {
		return fake.CreateStatsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createStatsReturns.result1
}

func (fake *FakeVRopsClient) CreateStatsCallCount() int {
	fake.createStatsMutex.RLock()
	defer fake.createStatsMutex.RUnlock()
	return len(fake.createStatsArgsForCall)
}

func (fake *FakeVRopsClient) CreateStatsArgsForCall(i int) (string, models.Stats) {
	fake.createStatsMutex.RLock()
	defer fake.createStatsMutex.RUnlock()
	return fake.createStatsArgsForCall[i].arg1, fake.createStatsArgsForCall[i].arg2
}

func (fake *FakeVRopsClient) CreateStatsReturns(result1 error) {
	fake.CreateStatsStub = nil
	fake.createStatsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVRopsClient) CreateStatsReturnsOnCall(i int, result1 error) {
	fake.CreateStatsStub = nil
	if fake.createStatsReturnsOnCall == nil {
		fake.createStatsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createStatsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVRopsClient) GetStatsForResource(arg1 string, arg2 string, arg3 string) (models.ListStatsResponseValuesStatListStats, error) {
	fake.getStatsForResourceMutex.Lock()
	ret, specificReturn := fake.getStatsForResourceReturnsOnCall[len(fake.getStatsForResourceArgsForCall)]
	fake.getStatsForResourceArgsForCall = append(fake.getStatsForResourceArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetStatsForResource", []interface{}{arg1, arg2, arg3})
	fake.getStatsForResourceMutex.Unlock()
	if fake.GetStatsForResourceStub != nil {
		return fake.GetStatsForResourceStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getStatsForResourceReturns.result1, fake.getStatsForResourceReturns.result2
}

func (fake *FakeVRopsClient) GetStatsForResourceCallCount() int {
	fake.getStatsForResourceMutex.RLock()
	defer fake.getStatsForResourceMutex.RUnlock()
	return len(fake.getStatsForResourceArgsForCall)
}

func (fake *FakeVRopsClient) GetStatsForResourceArgsForCall(i int) (string, string, string) {
	fake.getStatsForResourceMutex.RLock()
	defer fake.getStatsForResourceMutex.RUnlock()
	return fake.getStatsForResourceArgsForCall[i].arg1, fake.getStatsForResourceArgsForCall[i].arg2, fake.getStatsForResourceArgsForCall[i].arg3
}

func (fake *FakeVRopsClient) GetStatsForResourceReturns(result1 models.ListStatsResponseValuesStatListStats, result2 error) {
	fake.GetStatsForResourceStub = nil
	fake.getStatsForResourceReturns = struct {
		result1 models.ListStatsResponseValuesStatListStats
		result2 error
	}{result1, result2}
}

func (fake *FakeVRopsClient) GetStatsForResourceReturnsOnCall(i int, result1 models.ListStatsResponseValuesStatListStats, result2 error) {
	fake.GetStatsForResourceStub = nil
	if fake.getStatsForResourceReturnsOnCall == nil {
		fake.getStatsForResourceReturnsOnCall = make(map[int]struct {
			result1 models.ListStatsResponseValuesStatListStats
			result2 error
		})
	}
	fake.getStatsForResourceReturnsOnCall[i] = struct {
		result1 models.ListStatsResponseValuesStatListStats
		result2 error
	}{result1, result2}
}

func (fake *FakeVRopsClient) CreateResource(arg1 models.Resource) error {
	fake.createResourceMutex.Lock()
	ret, specificReturn := fake.createResourceReturnsOnCall[len(fake.createResourceArgsForCall)]
	fake.createResourceArgsForCall = append(fake.createResourceArgsForCall, struct {
		arg1 models.Resource
	}{arg1})
	fake.recordInvocation("CreateResource", []interface{}{arg1})
	fake.createResourceMutex.Unlock()
	if fake.CreateResourceStub != nil {
		return fake.CreateResourceStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createResourceReturns.result1
}

func (fake *FakeVRopsClient) CreateResourceCallCount() int {
	fake.createResourceMutex.RLock()
	defer fake.createResourceMutex.RUnlock()
	return len(fake.createResourceArgsForCall)
}

func (fake *FakeVRopsClient) CreateResourceArgsForCall(i int) models.Resource {
	fake.createResourceMutex.RLock()
	defer fake.createResourceMutex.RUnlock()
	return fake.createResourceArgsForCall[i].arg1
}

func (fake *FakeVRopsClient) CreateResourceReturns(result1 error) {
	fake.CreateResourceStub = nil
	fake.createResourceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVRopsClient) CreateResourceReturnsOnCall(i int, result1 error) {
	fake.CreateResourceStub = nil
	if fake.createResourceReturnsOnCall == nil {
		fake.createResourceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createResourceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVRopsClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.adapterKindsMutex.RLock()
	defer fake.adapterKindsMutex.RUnlock()
	fake.resourceKindsMutex.RLock()
	defer fake.resourceKindsMutex.RUnlock()
	fake.resourcesForAdapterKindMutex.RLock()
	defer fake.resourcesForAdapterKindMutex.RUnlock()
	fake.findResourceMutex.RLock()
	defer fake.findResourceMutex.RUnlock()
	fake.createStatsMutex.RLock()
	defer fake.createStatsMutex.RUnlock()
	fake.getStatsForResourceMutex.RLock()
	defer fake.getStatsForResourceMutex.RUnlock()
	fake.createResourceMutex.RLock()
	defer fake.createResourceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVRopsClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ clients.VRopsClientIntf = new(FakeVRopsClient)
