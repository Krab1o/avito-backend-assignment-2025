// Code generated by http://github.com/gojuno/minimock (v3.4.5). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/Krab1o/avito-backend-assignment-2025/internal/repository.MerchRepository -o merch_repository_mimimock.go -n MerchRepositoryMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	merchModel "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/merch/model"
	"github.com/gojuno/minimock/v3"
	"github.com/jackc/pgx/v5"
)

// MerchRepositoryMock implements mm_repository.MerchRepository
type MerchRepositoryMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcGetItem          func(ctx context.Context, tx pgx.Tx, itemTitle string) (mp1 *merchModel.Merch, err error)
	funcGetItemOrigin    string
	inspectFuncGetItem   func(ctx context.Context, tx pgx.Tx, itemTitle string)
	afterGetItemCounter  uint64
	beforeGetItemCounter uint64
	GetItemMock          mMerchRepositoryMockGetItem
}

// NewMerchRepositoryMock returns a mock for mm_repository.MerchRepository
func NewMerchRepositoryMock(t minimock.Tester) *MerchRepositoryMock {
	m := &MerchRepositoryMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetItemMock = mMerchRepositoryMockGetItem{mock: m}
	m.GetItemMock.callArgs = []*MerchRepositoryMockGetItemParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mMerchRepositoryMockGetItem struct {
	optional           bool
	mock               *MerchRepositoryMock
	defaultExpectation *MerchRepositoryMockGetItemExpectation
	expectations       []*MerchRepositoryMockGetItemExpectation

	callArgs []*MerchRepositoryMockGetItemParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// MerchRepositoryMockGetItemExpectation specifies expectation struct of the MerchRepository.GetItem
type MerchRepositoryMockGetItemExpectation struct {
	mock               *MerchRepositoryMock
	params             *MerchRepositoryMockGetItemParams
	paramPtrs          *MerchRepositoryMockGetItemParamPtrs
	expectationOrigins MerchRepositoryMockGetItemExpectationOrigins
	results            *MerchRepositoryMockGetItemResults
	returnOrigin       string
	Counter            uint64
}

// MerchRepositoryMockGetItemParams contains parameters of the MerchRepository.GetItem
type MerchRepositoryMockGetItemParams struct {
	ctx       context.Context
	tx        pgx.Tx
	itemTitle string
}

// MerchRepositoryMockGetItemParamPtrs contains pointers to parameters of the MerchRepository.GetItem
type MerchRepositoryMockGetItemParamPtrs struct {
	ctx       *context.Context
	tx        *pgx.Tx
	itemTitle *string
}

// MerchRepositoryMockGetItemResults contains results of the MerchRepository.GetItem
type MerchRepositoryMockGetItemResults struct {
	mp1 *merchModel.Merch
	err error
}

// MerchRepositoryMockGetItemOrigins contains origins of expectations of the MerchRepository.GetItem
type MerchRepositoryMockGetItemExpectationOrigins struct {
	origin          string
	originCtx       string
	originTx        string
	originItemTitle string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmGetItem *mMerchRepositoryMockGetItem) Optional() *mMerchRepositoryMockGetItem {
	mmGetItem.optional = true
	return mmGetItem
}

// Expect sets up expected params for MerchRepository.GetItem
func (mmGetItem *mMerchRepositoryMockGetItem) Expect(ctx context.Context, tx pgx.Tx, itemTitle string) *mMerchRepositoryMockGetItem {
	if mmGetItem.mock.funcGetItem != nil {
		mmGetItem.mock.t.Fatalf("MerchRepositoryMock.GetItem mock is already set by Set")
	}

	if mmGetItem.defaultExpectation == nil {
		mmGetItem.defaultExpectation = &MerchRepositoryMockGetItemExpectation{}
	}

	if mmGetItem.defaultExpectation.paramPtrs != nil {
		mmGetItem.mock.t.Fatalf("MerchRepositoryMock.GetItem mock is already set by ExpectParams functions")
	}

	mmGetItem.defaultExpectation.params = &MerchRepositoryMockGetItemParams{ctx, tx, itemTitle}
	mmGetItem.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmGetItem.expectations {
		if minimock.Equal(e.params, mmGetItem.defaultExpectation.params) {
			mmGetItem.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetItem.defaultExpectation.params)
		}
	}

	return mmGetItem
}

// ExpectCtxParam1 sets up expected param ctx for MerchRepository.GetItem
func (mmGetItem *mMerchRepositoryMockGetItem) ExpectCtxParam1(ctx context.Context) *mMerchRepositoryMockGetItem {
	if mmGetItem.mock.funcGetItem != nil {
		mmGetItem.mock.t.Fatalf("MerchRepositoryMock.GetItem mock is already set by Set")
	}

	if mmGetItem.defaultExpectation == nil {
		mmGetItem.defaultExpectation = &MerchRepositoryMockGetItemExpectation{}
	}

	if mmGetItem.defaultExpectation.params != nil {
		mmGetItem.mock.t.Fatalf("MerchRepositoryMock.GetItem mock is already set by Expect")
	}

	if mmGetItem.defaultExpectation.paramPtrs == nil {
		mmGetItem.defaultExpectation.paramPtrs = &MerchRepositoryMockGetItemParamPtrs{}
	}
	mmGetItem.defaultExpectation.paramPtrs.ctx = &ctx
	mmGetItem.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmGetItem
}

// ExpectTxParam2 sets up expected param tx for MerchRepository.GetItem
func (mmGetItem *mMerchRepositoryMockGetItem) ExpectTxParam2(tx pgx.Tx) *mMerchRepositoryMockGetItem {
	if mmGetItem.mock.funcGetItem != nil {
		mmGetItem.mock.t.Fatalf("MerchRepositoryMock.GetItem mock is already set by Set")
	}

	if mmGetItem.defaultExpectation == nil {
		mmGetItem.defaultExpectation = &MerchRepositoryMockGetItemExpectation{}
	}

	if mmGetItem.defaultExpectation.params != nil {
		mmGetItem.mock.t.Fatalf("MerchRepositoryMock.GetItem mock is already set by Expect")
	}

	if mmGetItem.defaultExpectation.paramPtrs == nil {
		mmGetItem.defaultExpectation.paramPtrs = &MerchRepositoryMockGetItemParamPtrs{}
	}
	mmGetItem.defaultExpectation.paramPtrs.tx = &tx
	mmGetItem.defaultExpectation.expectationOrigins.originTx = minimock.CallerInfo(1)

	return mmGetItem
}

// ExpectItemTitleParam3 sets up expected param itemTitle for MerchRepository.GetItem
func (mmGetItem *mMerchRepositoryMockGetItem) ExpectItemTitleParam3(itemTitle string) *mMerchRepositoryMockGetItem {
	if mmGetItem.mock.funcGetItem != nil {
		mmGetItem.mock.t.Fatalf("MerchRepositoryMock.GetItem mock is already set by Set")
	}

	if mmGetItem.defaultExpectation == nil {
		mmGetItem.defaultExpectation = &MerchRepositoryMockGetItemExpectation{}
	}

	if mmGetItem.defaultExpectation.params != nil {
		mmGetItem.mock.t.Fatalf("MerchRepositoryMock.GetItem mock is already set by Expect")
	}

	if mmGetItem.defaultExpectation.paramPtrs == nil {
		mmGetItem.defaultExpectation.paramPtrs = &MerchRepositoryMockGetItemParamPtrs{}
	}
	mmGetItem.defaultExpectation.paramPtrs.itemTitle = &itemTitle
	mmGetItem.defaultExpectation.expectationOrigins.originItemTitle = minimock.CallerInfo(1)

	return mmGetItem
}

// Inspect accepts an inspector function that has same arguments as the MerchRepository.GetItem
func (mmGetItem *mMerchRepositoryMockGetItem) Inspect(f func(ctx context.Context, tx pgx.Tx, itemTitle string)) *mMerchRepositoryMockGetItem {
	if mmGetItem.mock.inspectFuncGetItem != nil {
		mmGetItem.mock.t.Fatalf("Inspect function is already set for MerchRepositoryMock.GetItem")
	}

	mmGetItem.mock.inspectFuncGetItem = f

	return mmGetItem
}

// Return sets up results that will be returned by MerchRepository.GetItem
func (mmGetItem *mMerchRepositoryMockGetItem) Return(mp1 *merchModel.Merch, err error) *MerchRepositoryMock {
	if mmGetItem.mock.funcGetItem != nil {
		mmGetItem.mock.t.Fatalf("MerchRepositoryMock.GetItem mock is already set by Set")
	}

	if mmGetItem.defaultExpectation == nil {
		mmGetItem.defaultExpectation = &MerchRepositoryMockGetItemExpectation{mock: mmGetItem.mock}
	}
	mmGetItem.defaultExpectation.results = &MerchRepositoryMockGetItemResults{mp1, err}
	mmGetItem.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmGetItem.mock
}

// Set uses given function f to mock the MerchRepository.GetItem method
func (mmGetItem *mMerchRepositoryMockGetItem) Set(f func(ctx context.Context, tx pgx.Tx, itemTitle string) (mp1 *merchModel.Merch, err error)) *MerchRepositoryMock {
	if mmGetItem.defaultExpectation != nil {
		mmGetItem.mock.t.Fatalf("Default expectation is already set for the MerchRepository.GetItem method")
	}

	if len(mmGetItem.expectations) > 0 {
		mmGetItem.mock.t.Fatalf("Some expectations are already set for the MerchRepository.GetItem method")
	}

	mmGetItem.mock.funcGetItem = f
	mmGetItem.mock.funcGetItemOrigin = minimock.CallerInfo(1)
	return mmGetItem.mock
}

// When sets expectation for the MerchRepository.GetItem which will trigger the result defined by the following
// Then helper
func (mmGetItem *mMerchRepositoryMockGetItem) When(ctx context.Context, tx pgx.Tx, itemTitle string) *MerchRepositoryMockGetItemExpectation {
	if mmGetItem.mock.funcGetItem != nil {
		mmGetItem.mock.t.Fatalf("MerchRepositoryMock.GetItem mock is already set by Set")
	}

	expectation := &MerchRepositoryMockGetItemExpectation{
		mock:               mmGetItem.mock,
		params:             &MerchRepositoryMockGetItemParams{ctx, tx, itemTitle},
		expectationOrigins: MerchRepositoryMockGetItemExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmGetItem.expectations = append(mmGetItem.expectations, expectation)
	return expectation
}

// Then sets up MerchRepository.GetItem return parameters for the expectation previously defined by the When method
func (e *MerchRepositoryMockGetItemExpectation) Then(mp1 *merchModel.Merch, err error) *MerchRepositoryMock {
	e.results = &MerchRepositoryMockGetItemResults{mp1, err}
	return e.mock
}

// Times sets number of times MerchRepository.GetItem should be invoked
func (mmGetItem *mMerchRepositoryMockGetItem) Times(n uint64) *mMerchRepositoryMockGetItem {
	if n == 0 {
		mmGetItem.mock.t.Fatalf("Times of MerchRepositoryMock.GetItem mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmGetItem.expectedInvocations, n)
	mmGetItem.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmGetItem
}

func (mmGetItem *mMerchRepositoryMockGetItem) invocationsDone() bool {
	if len(mmGetItem.expectations) == 0 && mmGetItem.defaultExpectation == nil && mmGetItem.mock.funcGetItem == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmGetItem.mock.afterGetItemCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmGetItem.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// GetItem implements mm_repository.MerchRepository
func (mmGetItem *MerchRepositoryMock) GetItem(ctx context.Context, tx pgx.Tx, itemTitle string) (mp1 *merchModel.Merch, err error) {
	mm_atomic.AddUint64(&mmGetItem.beforeGetItemCounter, 1)
	defer mm_atomic.AddUint64(&mmGetItem.afterGetItemCounter, 1)

	mmGetItem.t.Helper()

	if mmGetItem.inspectFuncGetItem != nil {
		mmGetItem.inspectFuncGetItem(ctx, tx, itemTitle)
	}

	mm_params := MerchRepositoryMockGetItemParams{ctx, tx, itemTitle}

	// Record call args
	mmGetItem.GetItemMock.mutex.Lock()
	mmGetItem.GetItemMock.callArgs = append(mmGetItem.GetItemMock.callArgs, &mm_params)
	mmGetItem.GetItemMock.mutex.Unlock()

	for _, e := range mmGetItem.GetItemMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.mp1, e.results.err
		}
	}

	if mmGetItem.GetItemMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetItem.GetItemMock.defaultExpectation.Counter, 1)
		mm_want := mmGetItem.GetItemMock.defaultExpectation.params
		mm_want_ptrs := mmGetItem.GetItemMock.defaultExpectation.paramPtrs

		mm_got := MerchRepositoryMockGetItemParams{ctx, tx, itemTitle}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmGetItem.t.Errorf("MerchRepositoryMock.GetItem got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmGetItem.GetItemMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.tx != nil && !minimock.Equal(*mm_want_ptrs.tx, mm_got.tx) {
				mmGetItem.t.Errorf("MerchRepositoryMock.GetItem got unexpected parameter tx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmGetItem.GetItemMock.defaultExpectation.expectationOrigins.originTx, *mm_want_ptrs.tx, mm_got.tx, minimock.Diff(*mm_want_ptrs.tx, mm_got.tx))
			}

			if mm_want_ptrs.itemTitle != nil && !minimock.Equal(*mm_want_ptrs.itemTitle, mm_got.itemTitle) {
				mmGetItem.t.Errorf("MerchRepositoryMock.GetItem got unexpected parameter itemTitle, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmGetItem.GetItemMock.defaultExpectation.expectationOrigins.originItemTitle, *mm_want_ptrs.itemTitle, mm_got.itemTitle, minimock.Diff(*mm_want_ptrs.itemTitle, mm_got.itemTitle))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetItem.t.Errorf("MerchRepositoryMock.GetItem got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmGetItem.GetItemMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetItem.GetItemMock.defaultExpectation.results
		if mm_results == nil {
			mmGetItem.t.Fatal("No results are set for the MerchRepositoryMock.GetItem")
		}
		return (*mm_results).mp1, (*mm_results).err
	}
	if mmGetItem.funcGetItem != nil {
		return mmGetItem.funcGetItem(ctx, tx, itemTitle)
	}
	mmGetItem.t.Fatalf("Unexpected call to MerchRepositoryMock.GetItem. %v %v %v", ctx, tx, itemTitle)
	return
}

// GetItemAfterCounter returns a count of finished MerchRepositoryMock.GetItem invocations
func (mmGetItem *MerchRepositoryMock) GetItemAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetItem.afterGetItemCounter)
}

// GetItemBeforeCounter returns a count of MerchRepositoryMock.GetItem invocations
func (mmGetItem *MerchRepositoryMock) GetItemBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetItem.beforeGetItemCounter)
}

// Calls returns a list of arguments used in each call to MerchRepositoryMock.GetItem.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetItem *mMerchRepositoryMockGetItem) Calls() []*MerchRepositoryMockGetItemParams {
	mmGetItem.mutex.RLock()

	argCopy := make([]*MerchRepositoryMockGetItemParams, len(mmGetItem.callArgs))
	copy(argCopy, mmGetItem.callArgs)

	mmGetItem.mutex.RUnlock()

	return argCopy
}

// MinimockGetItemDone returns true if the count of the GetItem invocations corresponds
// the number of defined expectations
func (m *MerchRepositoryMock) MinimockGetItemDone() bool {
	if m.GetItemMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.GetItemMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.GetItemMock.invocationsDone()
}

// MinimockGetItemInspect logs each unmet expectation
func (m *MerchRepositoryMock) MinimockGetItemInspect() {
	for _, e := range m.GetItemMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MerchRepositoryMock.GetItem at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterGetItemCounter := mm_atomic.LoadUint64(&m.afterGetItemCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.GetItemMock.defaultExpectation != nil && afterGetItemCounter < 1 {
		if m.GetItemMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to MerchRepositoryMock.GetItem at\n%s", m.GetItemMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to MerchRepositoryMock.GetItem at\n%s with params: %#v", m.GetItemMock.defaultExpectation.expectationOrigins.origin, *m.GetItemMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetItem != nil && afterGetItemCounter < 1 {
		m.t.Errorf("Expected call to MerchRepositoryMock.GetItem at\n%s", m.funcGetItemOrigin)
	}

	if !m.GetItemMock.invocationsDone() && afterGetItemCounter > 0 {
		m.t.Errorf("Expected %d calls to MerchRepositoryMock.GetItem at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.GetItemMock.expectedInvocations), m.GetItemMock.expectedInvocationsOrigin, afterGetItemCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *MerchRepositoryMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockGetItemInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *MerchRepositoryMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *MerchRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetItemDone()
}
