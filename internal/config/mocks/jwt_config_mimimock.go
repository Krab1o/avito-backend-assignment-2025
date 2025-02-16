// Code generated by http://github.com/gojuno/minimock (v3.4.5). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/Krab1o/avito-backend-assignment-2025/internal/config.JWTConfig -o jwt_config_mimimock.go -n JWTConfigMock -p mocks

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// JWTConfigMock implements mm_config.JWTConfig
type JWTConfigMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcSecret          func() (ba1 []byte)
	funcSecretOrigin    string
	inspectFuncSecret   func()
	afterSecretCounter  uint64
	beforeSecretCounter uint64
	SecretMock          mJWTConfigMockSecret

	funcTimeout          func() (i1 int)
	funcTimeoutOrigin    string
	inspectFuncTimeout   func()
	afterTimeoutCounter  uint64
	beforeTimeoutCounter uint64
	TimeoutMock          mJWTConfigMockTimeout
}

// NewJWTConfigMock returns a mock for mm_config.JWTConfig
func NewJWTConfigMock(t minimock.Tester) *JWTConfigMock {
	m := &JWTConfigMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.SecretMock = mJWTConfigMockSecret{mock: m}

	m.TimeoutMock = mJWTConfigMockTimeout{mock: m}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mJWTConfigMockSecret struct {
	optional           bool
	mock               *JWTConfigMock
	defaultExpectation *JWTConfigMockSecretExpectation
	expectations       []*JWTConfigMockSecretExpectation

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// JWTConfigMockSecretExpectation specifies expectation struct of the JWTConfig.Secret
type JWTConfigMockSecretExpectation struct {
	mock *JWTConfigMock

	results      *JWTConfigMockSecretResults
	returnOrigin string
	Counter      uint64
}

// JWTConfigMockSecretResults contains results of the JWTConfig.Secret
type JWTConfigMockSecretResults struct {
	ba1 []byte
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmSecret *mJWTConfigMockSecret) Optional() *mJWTConfigMockSecret {
	mmSecret.optional = true
	return mmSecret
}

// Expect sets up expected params for JWTConfig.Secret
func (mmSecret *mJWTConfigMockSecret) Expect() *mJWTConfigMockSecret {
	if mmSecret.mock.funcSecret != nil {
		mmSecret.mock.t.Fatalf("JWTConfigMock.Secret mock is already set by Set")
	}

	if mmSecret.defaultExpectation == nil {
		mmSecret.defaultExpectation = &JWTConfigMockSecretExpectation{}
	}

	return mmSecret
}

// Inspect accepts an inspector function that has same arguments as the JWTConfig.Secret
func (mmSecret *mJWTConfigMockSecret) Inspect(f func()) *mJWTConfigMockSecret {
	if mmSecret.mock.inspectFuncSecret != nil {
		mmSecret.mock.t.Fatalf("Inspect function is already set for JWTConfigMock.Secret")
	}

	mmSecret.mock.inspectFuncSecret = f

	return mmSecret
}

// Return sets up results that will be returned by JWTConfig.Secret
func (mmSecret *mJWTConfigMockSecret) Return(ba1 []byte) *JWTConfigMock {
	if mmSecret.mock.funcSecret != nil {
		mmSecret.mock.t.Fatalf("JWTConfigMock.Secret mock is already set by Set")
	}

	if mmSecret.defaultExpectation == nil {
		mmSecret.defaultExpectation = &JWTConfigMockSecretExpectation{mock: mmSecret.mock}
	}
	mmSecret.defaultExpectation.results = &JWTConfigMockSecretResults{ba1}
	mmSecret.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmSecret.mock
}

// Set uses given function f to mock the JWTConfig.Secret method
func (mmSecret *mJWTConfigMockSecret) Set(f func() (ba1 []byte)) *JWTConfigMock {
	if mmSecret.defaultExpectation != nil {
		mmSecret.mock.t.Fatalf("Default expectation is already set for the JWTConfig.Secret method")
	}

	if len(mmSecret.expectations) > 0 {
		mmSecret.mock.t.Fatalf("Some expectations are already set for the JWTConfig.Secret method")
	}

	mmSecret.mock.funcSecret = f
	mmSecret.mock.funcSecretOrigin = minimock.CallerInfo(1)
	return mmSecret.mock
}

// Times sets number of times JWTConfig.Secret should be invoked
func (mmSecret *mJWTConfigMockSecret) Times(n uint64) *mJWTConfigMockSecret {
	if n == 0 {
		mmSecret.mock.t.Fatalf("Times of JWTConfigMock.Secret mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmSecret.expectedInvocations, n)
	mmSecret.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmSecret
}

func (mmSecret *mJWTConfigMockSecret) invocationsDone() bool {
	if len(mmSecret.expectations) == 0 && mmSecret.defaultExpectation == nil && mmSecret.mock.funcSecret == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmSecret.mock.afterSecretCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmSecret.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Secret implements mm_config.JWTConfig
func (mmSecret *JWTConfigMock) Secret() (ba1 []byte) {
	mm_atomic.AddUint64(&mmSecret.beforeSecretCounter, 1)
	defer mm_atomic.AddUint64(&mmSecret.afterSecretCounter, 1)

	mmSecret.t.Helper()

	if mmSecret.inspectFuncSecret != nil {
		mmSecret.inspectFuncSecret()
	}

	if mmSecret.SecretMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSecret.SecretMock.defaultExpectation.Counter, 1)

		mm_results := mmSecret.SecretMock.defaultExpectation.results
		if mm_results == nil {
			mmSecret.t.Fatal("No results are set for the JWTConfigMock.Secret")
		}
		return (*mm_results).ba1
	}
	if mmSecret.funcSecret != nil {
		return mmSecret.funcSecret()
	}
	mmSecret.t.Fatalf("Unexpected call to JWTConfigMock.Secret.")
	return
}

// SecretAfterCounter returns a count of finished JWTConfigMock.Secret invocations
func (mmSecret *JWTConfigMock) SecretAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSecret.afterSecretCounter)
}

// SecretBeforeCounter returns a count of JWTConfigMock.Secret invocations
func (mmSecret *JWTConfigMock) SecretBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSecret.beforeSecretCounter)
}

// MinimockSecretDone returns true if the count of the Secret invocations corresponds
// the number of defined expectations
func (m *JWTConfigMock) MinimockSecretDone() bool {
	if m.SecretMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.SecretMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.SecretMock.invocationsDone()
}

// MinimockSecretInspect logs each unmet expectation
func (m *JWTConfigMock) MinimockSecretInspect() {
	for _, e := range m.SecretMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to JWTConfigMock.Secret")
		}
	}

	afterSecretCounter := mm_atomic.LoadUint64(&m.afterSecretCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.SecretMock.defaultExpectation != nil && afterSecretCounter < 1 {
		m.t.Errorf("Expected call to JWTConfigMock.Secret at\n%s", m.SecretMock.defaultExpectation.returnOrigin)
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSecret != nil && afterSecretCounter < 1 {
		m.t.Errorf("Expected call to JWTConfigMock.Secret at\n%s", m.funcSecretOrigin)
	}

	if !m.SecretMock.invocationsDone() && afterSecretCounter > 0 {
		m.t.Errorf("Expected %d calls to JWTConfigMock.Secret at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.SecretMock.expectedInvocations), m.SecretMock.expectedInvocationsOrigin, afterSecretCounter)
	}
}

type mJWTConfigMockTimeout struct {
	optional           bool
	mock               *JWTConfigMock
	defaultExpectation *JWTConfigMockTimeoutExpectation
	expectations       []*JWTConfigMockTimeoutExpectation

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// JWTConfigMockTimeoutExpectation specifies expectation struct of the JWTConfig.Timeout
type JWTConfigMockTimeoutExpectation struct {
	mock *JWTConfigMock

	results      *JWTConfigMockTimeoutResults
	returnOrigin string
	Counter      uint64
}

// JWTConfigMockTimeoutResults contains results of the JWTConfig.Timeout
type JWTConfigMockTimeoutResults struct {
	i1 int
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmTimeout *mJWTConfigMockTimeout) Optional() *mJWTConfigMockTimeout {
	mmTimeout.optional = true
	return mmTimeout
}

// Expect sets up expected params for JWTConfig.Timeout
func (mmTimeout *mJWTConfigMockTimeout) Expect() *mJWTConfigMockTimeout {
	if mmTimeout.mock.funcTimeout != nil {
		mmTimeout.mock.t.Fatalf("JWTConfigMock.Timeout mock is already set by Set")
	}

	if mmTimeout.defaultExpectation == nil {
		mmTimeout.defaultExpectation = &JWTConfigMockTimeoutExpectation{}
	}

	return mmTimeout
}

// Inspect accepts an inspector function that has same arguments as the JWTConfig.Timeout
func (mmTimeout *mJWTConfigMockTimeout) Inspect(f func()) *mJWTConfigMockTimeout {
	if mmTimeout.mock.inspectFuncTimeout != nil {
		mmTimeout.mock.t.Fatalf("Inspect function is already set for JWTConfigMock.Timeout")
	}

	mmTimeout.mock.inspectFuncTimeout = f

	return mmTimeout
}

// Return sets up results that will be returned by JWTConfig.Timeout
func (mmTimeout *mJWTConfigMockTimeout) Return(i1 int) *JWTConfigMock {
	if mmTimeout.mock.funcTimeout != nil {
		mmTimeout.mock.t.Fatalf("JWTConfigMock.Timeout mock is already set by Set")
	}

	if mmTimeout.defaultExpectation == nil {
		mmTimeout.defaultExpectation = &JWTConfigMockTimeoutExpectation{mock: mmTimeout.mock}
	}
	mmTimeout.defaultExpectation.results = &JWTConfigMockTimeoutResults{i1}
	mmTimeout.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmTimeout.mock
}

// Set uses given function f to mock the JWTConfig.Timeout method
func (mmTimeout *mJWTConfigMockTimeout) Set(f func() (i1 int)) *JWTConfigMock {
	if mmTimeout.defaultExpectation != nil {
		mmTimeout.mock.t.Fatalf("Default expectation is already set for the JWTConfig.Timeout method")
	}

	if len(mmTimeout.expectations) > 0 {
		mmTimeout.mock.t.Fatalf("Some expectations are already set for the JWTConfig.Timeout method")
	}

	mmTimeout.mock.funcTimeout = f
	mmTimeout.mock.funcTimeoutOrigin = minimock.CallerInfo(1)
	return mmTimeout.mock
}

// Times sets number of times JWTConfig.Timeout should be invoked
func (mmTimeout *mJWTConfigMockTimeout) Times(n uint64) *mJWTConfigMockTimeout {
	if n == 0 {
		mmTimeout.mock.t.Fatalf("Times of JWTConfigMock.Timeout mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmTimeout.expectedInvocations, n)
	mmTimeout.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmTimeout
}

func (mmTimeout *mJWTConfigMockTimeout) invocationsDone() bool {
	if len(mmTimeout.expectations) == 0 && mmTimeout.defaultExpectation == nil && mmTimeout.mock.funcTimeout == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmTimeout.mock.afterTimeoutCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmTimeout.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Timeout implements mm_config.JWTConfig
func (mmTimeout *JWTConfigMock) Timeout() (i1 int) {
	mm_atomic.AddUint64(&mmTimeout.beforeTimeoutCounter, 1)
	defer mm_atomic.AddUint64(&mmTimeout.afterTimeoutCounter, 1)

	mmTimeout.t.Helper()

	if mmTimeout.inspectFuncTimeout != nil {
		mmTimeout.inspectFuncTimeout()
	}

	if mmTimeout.TimeoutMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmTimeout.TimeoutMock.defaultExpectation.Counter, 1)

		mm_results := mmTimeout.TimeoutMock.defaultExpectation.results
		if mm_results == nil {
			mmTimeout.t.Fatal("No results are set for the JWTConfigMock.Timeout")
		}
		return (*mm_results).i1
	}
	if mmTimeout.funcTimeout != nil {
		return mmTimeout.funcTimeout()
	}
	mmTimeout.t.Fatalf("Unexpected call to JWTConfigMock.Timeout.")
	return
}

// TimeoutAfterCounter returns a count of finished JWTConfigMock.Timeout invocations
func (mmTimeout *JWTConfigMock) TimeoutAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmTimeout.afterTimeoutCounter)
}

// TimeoutBeforeCounter returns a count of JWTConfigMock.Timeout invocations
func (mmTimeout *JWTConfigMock) TimeoutBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmTimeout.beforeTimeoutCounter)
}

// MinimockTimeoutDone returns true if the count of the Timeout invocations corresponds
// the number of defined expectations
func (m *JWTConfigMock) MinimockTimeoutDone() bool {
	if m.TimeoutMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.TimeoutMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.TimeoutMock.invocationsDone()
}

// MinimockTimeoutInspect logs each unmet expectation
func (m *JWTConfigMock) MinimockTimeoutInspect() {
	for _, e := range m.TimeoutMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to JWTConfigMock.Timeout")
		}
	}

	afterTimeoutCounter := mm_atomic.LoadUint64(&m.afterTimeoutCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.TimeoutMock.defaultExpectation != nil && afterTimeoutCounter < 1 {
		m.t.Errorf("Expected call to JWTConfigMock.Timeout at\n%s", m.TimeoutMock.defaultExpectation.returnOrigin)
	}
	// if func was set then invocations count should be greater than zero
	if m.funcTimeout != nil && afterTimeoutCounter < 1 {
		m.t.Errorf("Expected call to JWTConfigMock.Timeout at\n%s", m.funcTimeoutOrigin)
	}

	if !m.TimeoutMock.invocationsDone() && afterTimeoutCounter > 0 {
		m.t.Errorf("Expected %d calls to JWTConfigMock.Timeout at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.TimeoutMock.expectedInvocations), m.TimeoutMock.expectedInvocationsOrigin, afterTimeoutCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *JWTConfigMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockSecretInspect()

			m.MinimockTimeoutInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *JWTConfigMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *JWTConfigMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockSecretDone() &&
		m.MinimockTimeoutDone()
}
