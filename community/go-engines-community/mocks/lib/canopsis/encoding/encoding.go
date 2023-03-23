// Code generated by MockGen. DO NOT EDIT.
// Source: git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/encoding (interfaces: Encoder,Decoder)

// Package mock_encoding is a generated GoMock package.
package mock_encoding

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEncoder is a mock of Encoder interface.
type MockEncoder struct {
	ctrl     *gomock.Controller
	recorder *MockEncoderMockRecorder
}

// MockEncoderMockRecorder is the mock recorder for MockEncoder.
type MockEncoderMockRecorder struct {
	mock *MockEncoder
}

// NewMockEncoder creates a new mock instance.
func NewMockEncoder(ctrl *gomock.Controller) *MockEncoder {
	mock := &MockEncoder{ctrl: ctrl}
	mock.recorder = &MockEncoderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEncoder) EXPECT() *MockEncoderMockRecorder {
	return m.recorder
}

// Encode mocks base method.
func (m *MockEncoder) Encode(arg0 interface{}) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encode", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encode indicates an expected call of Encode.
func (mr *MockEncoderMockRecorder) Encode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encode", reflect.TypeOf((*MockEncoder)(nil).Encode), arg0)
}

// MockDecoder is a mock of Decoder interface.
type MockDecoder struct {
	ctrl     *gomock.Controller
	recorder *MockDecoderMockRecorder
}

// MockDecoderMockRecorder is the mock recorder for MockDecoder.
type MockDecoderMockRecorder struct {
	mock *MockDecoder
}

// NewMockDecoder creates a new mock instance.
func NewMockDecoder(ctrl *gomock.Controller) *MockDecoder {
	mock := &MockDecoder{ctrl: ctrl}
	mock.recorder = &MockDecoderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDecoder) EXPECT() *MockDecoderMockRecorder {
	return m.recorder
}

// Decode mocks base method.
func (m *MockDecoder) Decode(arg0 []byte, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Decode indicates an expected call of Decode.
func (mr *MockDecoderMockRecorder) Decode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockDecoder)(nil).Decode), arg0, arg1)
}