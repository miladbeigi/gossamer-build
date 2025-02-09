// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ChainSafe/gossamer/dot/digest (interfaces: GrandpaState)

// Package digest is a generated GoMock package.
package digest

import (
	reflect "reflect"

	types "github.com/ChainSafe/gossamer/dot/types"
	scale "github.com/ChainSafe/gossamer/pkg/scale"
	gomock "github.com/golang/mock/gomock"
)

// MockGrandpaState is a mock of GrandpaState interface.
type MockGrandpaState struct {
	ctrl     *gomock.Controller
	recorder *MockGrandpaStateMockRecorder
}

// MockGrandpaStateMockRecorder is the mock recorder for MockGrandpaState.
type MockGrandpaStateMockRecorder struct {
	mock *MockGrandpaState
}

// NewMockGrandpaState creates a new mock instance.
func NewMockGrandpaState(ctrl *gomock.Controller) *MockGrandpaState {
	mock := &MockGrandpaState{ctrl: ctrl}
	mock.recorder = &MockGrandpaStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGrandpaState) EXPECT() *MockGrandpaStateMockRecorder {
	return m.recorder
}

// ApplyScheduledChanges mocks base method.
func (m *MockGrandpaState) ApplyScheduledChanges(arg0 *types.Header) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyScheduledChanges", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyScheduledChanges indicates an expected call of ApplyScheduledChanges.
func (mr *MockGrandpaStateMockRecorder) ApplyScheduledChanges(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyScheduledChanges", reflect.TypeOf((*MockGrandpaState)(nil).ApplyScheduledChanges), arg0)
}

// HandleGRANDPADigest mocks base method.
func (m *MockGrandpaState) HandleGRANDPADigest(arg0 *types.Header, arg1 scale.VaryingDataType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleGRANDPADigest", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleGRANDPADigest indicates an expected call of HandleGRANDPADigest.
func (mr *MockGrandpaStateMockRecorder) HandleGRANDPADigest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleGRANDPADigest", reflect.TypeOf((*MockGrandpaState)(nil).HandleGRANDPADigest), arg0, arg1)
}
