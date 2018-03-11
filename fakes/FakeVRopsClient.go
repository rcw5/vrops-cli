// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/rcw5/vrops-cli/clients"
	"github.com/rcw5/vrops-cli/models"
)

type FakeVRopsClient struct {
	AdapterKindsStub        func() ([]models.AdapterKind, error)
	adapterKindsMutex       sync.RWMutex
	adapterKindsArgsForCall []struct{}
	adapterKindsReturns     struct {
		result1 []models.AdapterKind
		result2 error
	}
	adapterKindsReturnsOnCall map[int]struct {
		result1 []models.AdapterKind
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
	ResourcesForAdapterKindStub        func(string) ([]models.Resource, error)
	resourcesForAdapterKindMutex       sync.RWMutex
	resourcesForAdapterKindArgsForCall []struct {
		arg1 string
	}
	resourcesForAdapterKindReturns struct {
		result1 []models.Resource
		result2 error
	}
	resourcesForAdapterKindReturnsOnCall map[int]struct {
		result1 []models.Resource
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVRopsClient) AdapterKinds() ([]models.AdapterKind, error) {
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

func (fake *FakeVRopsClient) AdapterKindsReturns(result1 []models.AdapterKind, result2 error) {
	fake.AdapterKindsStub = nil
	fake.adapterKindsReturns = struct {
		result1 []models.AdapterKind
		result2 error
	}{result1, result2}
}

func (fake *FakeVRopsClient) AdapterKindsReturnsOnCall(i int, result1 []models.AdapterKind, result2 error) {
	fake.AdapterKindsStub = nil
	if fake.adapterKindsReturnsOnCall == nil {
		fake.adapterKindsReturnsOnCall = make(map[int]struct {
			result1 []models.AdapterKind
			result2 error
		})
	}
	fake.adapterKindsReturnsOnCall[i] = struct {
		result1 []models.AdapterKind
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

func (fake *FakeVRopsClient) ResourcesForAdapterKind(arg1 string) ([]models.Resource, error) {
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

func (fake *FakeVRopsClient) ResourcesForAdapterKindReturns(result1 []models.Resource, result2 error) {
	fake.ResourcesForAdapterKindStub = nil
	fake.resourcesForAdapterKindReturns = struct {
		result1 []models.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeVRopsClient) ResourcesForAdapterKindReturnsOnCall(i int, result1 []models.Resource, result2 error) {
	fake.ResourcesForAdapterKindStub = nil
	if fake.resourcesForAdapterKindReturnsOnCall == nil {
		fake.resourcesForAdapterKindReturnsOnCall = make(map[int]struct {
			result1 []models.Resource
			result2 error
		})
	}
	fake.resourcesForAdapterKindReturnsOnCall[i] = struct {
		result1 []models.Resource
		result2 error
	}{result1, result2}
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