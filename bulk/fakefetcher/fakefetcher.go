// This file was generated by counterfeiter
package fakefetcher

import (
	. "github.com/cloudfoundry-incubator/nsync/bulk"
	"github.com/cloudfoundry-incubator/runtime-schema/models"

	"sync"
)

type FakeFetcher struct {
	FetchStub        func(chan<- models.DesiredLRP) error
	fetchMutex       sync.RWMutex
	fetchArgsForCall []struct {
		arg1 chan<- models.DesiredLRP
	}
	fetchReturns struct {
		result1 error
	}
}

func (fake *FakeFetcher) Fetch(arg1 chan<- models.DesiredLRP) error {
	fake.fetchMutex.Lock()
	defer fake.fetchMutex.Unlock()
	fake.fetchArgsForCall = append(fake.fetchArgsForCall, struct {
		arg1 chan<- models.DesiredLRP
	}{arg1})
	if fake.FetchStub != nil {
		return fake.FetchStub(arg1)
	} else {
		return fake.fetchReturns.result1
	}
}

func (fake *FakeFetcher) FetchCallCount() int {
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	return len(fake.fetchArgsForCall)
}

func (fake *FakeFetcher) FetchArgsForCall(i int) chan<- models.DesiredLRP {
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	return fake.fetchArgsForCall[i].arg1
}

func (fake *FakeFetcher) FetchReturns(result1 error) {
	fake.fetchReturns = struct {
		result1 error
	}{result1}
}

var _ Fetcher = new(FakeFetcher)
