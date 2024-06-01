// Code generated by MockGen. DO NOT EDIT.
// Source: encoding.go

// Package mockEncoding is a generated GoMock package.
package mockEncoding

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEncoding is a mock of Encoding interface.
type MockEncoding struct {
	ctrl     *gomock.Controller
	recorder *MockEncodingMockRecorder
}

// MockEncodingMockRecorder is the mock recorder for MockEncoding.
type MockEncodingMockRecorder struct {
	mock *MockEncoding
}

// NewMockEncoding creates a new mock instance.
func NewMockEncoding(ctrl *gomock.Controller) *MockEncoding {
	mock := &MockEncoding{ctrl: ctrl}
	mock.recorder = &MockEncodingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEncoding) EXPECT() *MockEncodingMockRecorder {
	return m.recorder
}

// Compare mocks base method.
func (m *MockEncoding) Compare(encoded, raw string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Compare", encoded, raw)
	ret0, _ := ret[0].(error)
	return ret0
}

// Compare indicates an expected call of Compare.
func (mr *MockEncodingMockRecorder) Compare(encoded, raw interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compare", reflect.TypeOf((*MockEncoding)(nil).Compare), encoded, raw)
}

// Decode mocks base method.
func (m *MockEncoding) Decode(encoded string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", encoded)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decode indicates an expected call of Decode.
func (mr *MockEncodingMockRecorder) Decode(encoded interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockEncoding)(nil).Decode), encoded)
}

// Encode mocks base method.
func (m *MockEncoding) Encode(raw string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encode", raw)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encode indicates an expected call of Encode.
func (mr *MockEncodingMockRecorder) Encode(raw interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encode", reflect.TypeOf((*MockEncoding)(nil).Encode), raw)
}
