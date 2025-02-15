// Code generated by http://github.com/gojuno/minimock (v3.4.5). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/Krab1o/avito-backend-assignment-2025/internal/service.TransactionService -o transaction_service_mimimock.go -n TransactionServiceMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	transactionModel "github.com/Krab1o/avito-backend-assignment-2025/internal/service/transaction/model"
	"github.com/gojuno/minimock/v3"
)

// TransactionServiceMock implements mm_service.TransactionService
type TransactionServiceMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcSendCoin          func(ctx context.Context, tp1 *transactionModel.Transaction) (err error)
	funcSendCoinOrigin    string
	inspectFuncSendCoin   func(ctx context.Context, tp1 *transactionModel.Transaction)
	afterSendCoinCounter  uint64
	beforeSendCoinCounter uint64
	SendCoinMock          mTransactionServiceMockSendCoin
}

// NewTransactionServiceMock returns a mock for mm_service.TransactionService
func NewTransactionServiceMock(t minimock.Tester) *TransactionServiceMock {
	m := &TransactionServiceMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.SendCoinMock = mTransactionServiceMockSendCoin{mock: m}
	m.SendCoinMock.callArgs = []*TransactionServiceMockSendCoinParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mTransactionServiceMockSendCoin struct {
	optional           bool
	mock               *TransactionServiceMock
	defaultExpectation *TransactionServiceMockSendCoinExpectation
	expectations       []*TransactionServiceMockSendCoinExpectation

	callArgs []*TransactionServiceMockSendCoinParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// TransactionServiceMockSendCoinExpectation specifies expectation struct of the TransactionService.SendCoin
type TransactionServiceMockSendCoinExpectation struct {
	mock               *TransactionServiceMock
	params             *TransactionServiceMockSendCoinParams
	paramPtrs          *TransactionServiceMockSendCoinParamPtrs
	expectationOrigins TransactionServiceMockSendCoinExpectationOrigins
	results            *TransactionServiceMockSendCoinResults
	returnOrigin       string
	Counter            uint64
}

// TransactionServiceMockSendCoinParams contains parameters of the TransactionService.SendCoin
type TransactionServiceMockSendCoinParams struct {
	ctx context.Context
	tp1 *transactionModel.Transaction
}

// TransactionServiceMockSendCoinParamPtrs contains pointers to parameters of the TransactionService.SendCoin
type TransactionServiceMockSendCoinParamPtrs struct {
	ctx *context.Context
	tp1 **transactionModel.Transaction
}

// TransactionServiceMockSendCoinResults contains results of the TransactionService.SendCoin
type TransactionServiceMockSendCoinResults struct {
	err error
}

// TransactionServiceMockSendCoinOrigins contains origins of expectations of the TransactionService.SendCoin
type TransactionServiceMockSendCoinExpectationOrigins struct {
	origin    string
	originCtx string
	originTp1 string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmSendCoin *mTransactionServiceMockSendCoin) Optional() *mTransactionServiceMockSendCoin {
	mmSendCoin.optional = true
	return mmSendCoin
}

// Expect sets up expected params for TransactionService.SendCoin
func (mmSendCoin *mTransactionServiceMockSendCoin) Expect(ctx context.Context, tp1 *transactionModel.Transaction) *mTransactionServiceMockSendCoin {
	if mmSendCoin.mock.funcSendCoin != nil {
		mmSendCoin.mock.t.Fatalf("TransactionServiceMock.SendCoin mock is already set by Set")
	}

	if mmSendCoin.defaultExpectation == nil {
		mmSendCoin.defaultExpectation = &TransactionServiceMockSendCoinExpectation{}
	}

	if mmSendCoin.defaultExpectation.paramPtrs != nil {
		mmSendCoin.mock.t.Fatalf("TransactionServiceMock.SendCoin mock is already set by ExpectParams functions")
	}

	mmSendCoin.defaultExpectation.params = &TransactionServiceMockSendCoinParams{ctx, tp1}
	mmSendCoin.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmSendCoin.expectations {
		if minimock.Equal(e.params, mmSendCoin.defaultExpectation.params) {
			mmSendCoin.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSendCoin.defaultExpectation.params)
		}
	}

	return mmSendCoin
}

// ExpectCtxParam1 sets up expected param ctx for TransactionService.SendCoin
func (mmSendCoin *mTransactionServiceMockSendCoin) ExpectCtxParam1(ctx context.Context) *mTransactionServiceMockSendCoin {
	if mmSendCoin.mock.funcSendCoin != nil {
		mmSendCoin.mock.t.Fatalf("TransactionServiceMock.SendCoin mock is already set by Set")
	}

	if mmSendCoin.defaultExpectation == nil {
		mmSendCoin.defaultExpectation = &TransactionServiceMockSendCoinExpectation{}
	}

	if mmSendCoin.defaultExpectation.params != nil {
		mmSendCoin.mock.t.Fatalf("TransactionServiceMock.SendCoin mock is already set by Expect")
	}

	if mmSendCoin.defaultExpectation.paramPtrs == nil {
		mmSendCoin.defaultExpectation.paramPtrs = &TransactionServiceMockSendCoinParamPtrs{}
	}
	mmSendCoin.defaultExpectation.paramPtrs.ctx = &ctx
	mmSendCoin.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmSendCoin
}

// ExpectTp1Param2 sets up expected param tp1 for TransactionService.SendCoin
func (mmSendCoin *mTransactionServiceMockSendCoin) ExpectTp1Param2(tp1 *transactionModel.Transaction) *mTransactionServiceMockSendCoin {
	if mmSendCoin.mock.funcSendCoin != nil {
		mmSendCoin.mock.t.Fatalf("TransactionServiceMock.SendCoin mock is already set by Set")
	}

	if mmSendCoin.defaultExpectation == nil {
		mmSendCoin.defaultExpectation = &TransactionServiceMockSendCoinExpectation{}
	}

	if mmSendCoin.defaultExpectation.params != nil {
		mmSendCoin.mock.t.Fatalf("TransactionServiceMock.SendCoin mock is already set by Expect")
	}

	if mmSendCoin.defaultExpectation.paramPtrs == nil {
		mmSendCoin.defaultExpectation.paramPtrs = &TransactionServiceMockSendCoinParamPtrs{}
	}
	mmSendCoin.defaultExpectation.paramPtrs.tp1 = &tp1
	mmSendCoin.defaultExpectation.expectationOrigins.originTp1 = minimock.CallerInfo(1)

	return mmSendCoin
}

// Inspect accepts an inspector function that has same arguments as the TransactionService.SendCoin
func (mmSendCoin *mTransactionServiceMockSendCoin) Inspect(f func(ctx context.Context, tp1 *transactionModel.Transaction)) *mTransactionServiceMockSendCoin {
	if mmSendCoin.mock.inspectFuncSendCoin != nil {
		mmSendCoin.mock.t.Fatalf("Inspect function is already set for TransactionServiceMock.SendCoin")
	}

	mmSendCoin.mock.inspectFuncSendCoin = f

	return mmSendCoin
}

// Return sets up results that will be returned by TransactionService.SendCoin
func (mmSendCoin *mTransactionServiceMockSendCoin) Return(err error) *TransactionServiceMock {
	if mmSendCoin.mock.funcSendCoin != nil {
		mmSendCoin.mock.t.Fatalf("TransactionServiceMock.SendCoin mock is already set by Set")
	}

	if mmSendCoin.defaultExpectation == nil {
		mmSendCoin.defaultExpectation = &TransactionServiceMockSendCoinExpectation{mock: mmSendCoin.mock}
	}
	mmSendCoin.defaultExpectation.results = &TransactionServiceMockSendCoinResults{err}
	mmSendCoin.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmSendCoin.mock
}

// Set uses given function f to mock the TransactionService.SendCoin method
func (mmSendCoin *mTransactionServiceMockSendCoin) Set(f func(ctx context.Context, tp1 *transactionModel.Transaction) (err error)) *TransactionServiceMock {
	if mmSendCoin.defaultExpectation != nil {
		mmSendCoin.mock.t.Fatalf("Default expectation is already set for the TransactionService.SendCoin method")
	}

	if len(mmSendCoin.expectations) > 0 {
		mmSendCoin.mock.t.Fatalf("Some expectations are already set for the TransactionService.SendCoin method")
	}

	mmSendCoin.mock.funcSendCoin = f
	mmSendCoin.mock.funcSendCoinOrigin = minimock.CallerInfo(1)
	return mmSendCoin.mock
}

// When sets expectation for the TransactionService.SendCoin which will trigger the result defined by the following
// Then helper
func (mmSendCoin *mTransactionServiceMockSendCoin) When(ctx context.Context, tp1 *transactionModel.Transaction) *TransactionServiceMockSendCoinExpectation {
	if mmSendCoin.mock.funcSendCoin != nil {
		mmSendCoin.mock.t.Fatalf("TransactionServiceMock.SendCoin mock is already set by Set")
	}

	expectation := &TransactionServiceMockSendCoinExpectation{
		mock:               mmSendCoin.mock,
		params:             &TransactionServiceMockSendCoinParams{ctx, tp1},
		expectationOrigins: TransactionServiceMockSendCoinExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmSendCoin.expectations = append(mmSendCoin.expectations, expectation)
	return expectation
}

// Then sets up TransactionService.SendCoin return parameters for the expectation previously defined by the When method
func (e *TransactionServiceMockSendCoinExpectation) Then(err error) *TransactionServiceMock {
	e.results = &TransactionServiceMockSendCoinResults{err}
	return e.mock
}

// Times sets number of times TransactionService.SendCoin should be invoked
func (mmSendCoin *mTransactionServiceMockSendCoin) Times(n uint64) *mTransactionServiceMockSendCoin {
	if n == 0 {
		mmSendCoin.mock.t.Fatalf("Times of TransactionServiceMock.SendCoin mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmSendCoin.expectedInvocations, n)
	mmSendCoin.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmSendCoin
}

func (mmSendCoin *mTransactionServiceMockSendCoin) invocationsDone() bool {
	if len(mmSendCoin.expectations) == 0 && mmSendCoin.defaultExpectation == nil && mmSendCoin.mock.funcSendCoin == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmSendCoin.mock.afterSendCoinCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmSendCoin.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// SendCoin implements mm_service.TransactionService
func (mmSendCoin *TransactionServiceMock) SendCoin(ctx context.Context, tp1 *transactionModel.Transaction) (err error) {
	mm_atomic.AddUint64(&mmSendCoin.beforeSendCoinCounter, 1)
	defer mm_atomic.AddUint64(&mmSendCoin.afterSendCoinCounter, 1)

	mmSendCoin.t.Helper()

	if mmSendCoin.inspectFuncSendCoin != nil {
		mmSendCoin.inspectFuncSendCoin(ctx, tp1)
	}

	mm_params := TransactionServiceMockSendCoinParams{ctx, tp1}

	// Record call args
	mmSendCoin.SendCoinMock.mutex.Lock()
	mmSendCoin.SendCoinMock.callArgs = append(mmSendCoin.SendCoinMock.callArgs, &mm_params)
	mmSendCoin.SendCoinMock.mutex.Unlock()

	for _, e := range mmSendCoin.SendCoinMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmSendCoin.SendCoinMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSendCoin.SendCoinMock.defaultExpectation.Counter, 1)
		mm_want := mmSendCoin.SendCoinMock.defaultExpectation.params
		mm_want_ptrs := mmSendCoin.SendCoinMock.defaultExpectation.paramPtrs

		mm_got := TransactionServiceMockSendCoinParams{ctx, tp1}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmSendCoin.t.Errorf("TransactionServiceMock.SendCoin got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmSendCoin.SendCoinMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.tp1 != nil && !minimock.Equal(*mm_want_ptrs.tp1, mm_got.tp1) {
				mmSendCoin.t.Errorf("TransactionServiceMock.SendCoin got unexpected parameter tp1, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmSendCoin.SendCoinMock.defaultExpectation.expectationOrigins.originTp1, *mm_want_ptrs.tp1, mm_got.tp1, minimock.Diff(*mm_want_ptrs.tp1, mm_got.tp1))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSendCoin.t.Errorf("TransactionServiceMock.SendCoin got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmSendCoin.SendCoinMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSendCoin.SendCoinMock.defaultExpectation.results
		if mm_results == nil {
			mmSendCoin.t.Fatal("No results are set for the TransactionServiceMock.SendCoin")
		}
		return (*mm_results).err
	}
	if mmSendCoin.funcSendCoin != nil {
		return mmSendCoin.funcSendCoin(ctx, tp1)
	}
	mmSendCoin.t.Fatalf("Unexpected call to TransactionServiceMock.SendCoin. %v %v", ctx, tp1)
	return
}

// SendCoinAfterCounter returns a count of finished TransactionServiceMock.SendCoin invocations
func (mmSendCoin *TransactionServiceMock) SendCoinAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendCoin.afterSendCoinCounter)
}

// SendCoinBeforeCounter returns a count of TransactionServiceMock.SendCoin invocations
func (mmSendCoin *TransactionServiceMock) SendCoinBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendCoin.beforeSendCoinCounter)
}

// Calls returns a list of arguments used in each call to TransactionServiceMock.SendCoin.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSendCoin *mTransactionServiceMockSendCoin) Calls() []*TransactionServiceMockSendCoinParams {
	mmSendCoin.mutex.RLock()

	argCopy := make([]*TransactionServiceMockSendCoinParams, len(mmSendCoin.callArgs))
	copy(argCopy, mmSendCoin.callArgs)

	mmSendCoin.mutex.RUnlock()

	return argCopy
}

// MinimockSendCoinDone returns true if the count of the SendCoin invocations corresponds
// the number of defined expectations
func (m *TransactionServiceMock) MinimockSendCoinDone() bool {
	if m.SendCoinMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.SendCoinMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.SendCoinMock.invocationsDone()
}

// MinimockSendCoinInspect logs each unmet expectation
func (m *TransactionServiceMock) MinimockSendCoinInspect() {
	for _, e := range m.SendCoinMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to TransactionServiceMock.SendCoin at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterSendCoinCounter := mm_atomic.LoadUint64(&m.afterSendCoinCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.SendCoinMock.defaultExpectation != nil && afterSendCoinCounter < 1 {
		if m.SendCoinMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to TransactionServiceMock.SendCoin at\n%s", m.SendCoinMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to TransactionServiceMock.SendCoin at\n%s with params: %#v", m.SendCoinMock.defaultExpectation.expectationOrigins.origin, *m.SendCoinMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSendCoin != nil && afterSendCoinCounter < 1 {
		m.t.Errorf("Expected call to TransactionServiceMock.SendCoin at\n%s", m.funcSendCoinOrigin)
	}

	if !m.SendCoinMock.invocationsDone() && afterSendCoinCounter > 0 {
		m.t.Errorf("Expected %d calls to TransactionServiceMock.SendCoin at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.SendCoinMock.expectedInvocations), m.SendCoinMock.expectedInvocationsOrigin, afterSendCoinCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *TransactionServiceMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockSendCoinInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *TransactionServiceMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *TransactionServiceMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockSendCoinDone()
}
