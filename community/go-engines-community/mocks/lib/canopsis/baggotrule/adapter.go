// Code generated by MockGen. DO NOT EDIT.
// Source: git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/baggotrule (interfaces: Adapter)

// Package mock_baggotrule is a generated GoMock package.
package mock_baggotrule

import (
	context "context"
	reflect "reflect"

	baggotrule "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/baggotrule"
	gomock "github.com/golang/mock/gomock"
)

// MockAdapter is a mock of Adapter interface.
type MockAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockAdapterMockRecorder
}

// MockAdapterMockRecorder is the mock recorder for MockAdapter.
type MockAdapterMockRecorder struct {
	mock *MockAdapter
}

// NewMockAdapter creates a new mock instance.
func NewMockAdapter(ctrl *gomock.Controller) *MockAdapter {
	mock := &MockAdapter{ctrl: ctrl}
	mock.recorder = &MockAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdapter) EXPECT() *MockAdapterMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockAdapter) Get(arg0 context.Context) ([]baggotrule.Rule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].([]baggotrule.Rule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAdapterMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAdapter)(nil).Get), arg0)
}
