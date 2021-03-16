// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/JosiahWitt/ensure-cli/internal/fswrite (interfaces: FSWriteIface)

// Package mock_fswrite is a generated GoMock package.
package mock_fswrite

import (
	os "os"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFSWriteIface is a mock of FSWriteIface interface.
type MockFSWriteIface struct {
	ctrl     *gomock.Controller
	recorder *MockFSWriteIfaceMockRecorder
}

// MockFSWriteIfaceMockRecorder is the mock recorder for MockFSWriteIface.
type MockFSWriteIfaceMockRecorder struct {
	mock *MockFSWriteIface
}

// NewMockFSWriteIface creates a new mock instance.
func NewMockFSWriteIface(ctrl *gomock.Controller) *MockFSWriteIface {
	mock := &MockFSWriteIface{ctrl: ctrl}
	mock.recorder = &MockFSWriteIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFSWriteIface) EXPECT() *MockFSWriteIfaceMockRecorder {
	return m.recorder
}

// GlobRemoveAll mocks base method.
func (m *MockFSWriteIface) GlobRemoveAll(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GlobRemoveAll", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// GlobRemoveAll indicates an expected call of GlobRemoveAll.
func (mr *MockFSWriteIfaceMockRecorder) GlobRemoveAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GlobRemoveAll", reflect.TypeOf((*MockFSWriteIface)(nil).GlobRemoveAll), arg0)
}

// ListRecursive mocks base method.
func (m *MockFSWriteIface) ListRecursive(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRecursive", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRecursive indicates an expected call of ListRecursive.
func (mr *MockFSWriteIfaceMockRecorder) ListRecursive(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRecursive", reflect.TypeOf((*MockFSWriteIface)(nil).ListRecursive), arg0)
}

// MkdirAll mocks base method.
func (m *MockFSWriteIface) MkdirAll(arg0 string, arg1 os.FileMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MkdirAll", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MkdirAll indicates an expected call of MkdirAll.
func (mr *MockFSWriteIfaceMockRecorder) MkdirAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MkdirAll", reflect.TypeOf((*MockFSWriteIface)(nil).MkdirAll), arg0, arg1)
}

// RemoveAll mocks base method.
func (m *MockFSWriteIface) RemoveAll(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAll", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAll indicates an expected call of RemoveAll.
func (mr *MockFSWriteIfaceMockRecorder) RemoveAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAll", reflect.TypeOf((*MockFSWriteIface)(nil).RemoveAll), arg0)
}

// WriteFile mocks base method.
func (m *MockFSWriteIface) WriteFile(arg0, arg1 string, arg2 os.FileMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteFile", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteFile indicates an expected call of WriteFile.
func (mr *MockFSWriteIfaceMockRecorder) WriteFile(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteFile", reflect.TypeOf((*MockFSWriteIface)(nil).WriteFile), arg0, arg1, arg2)
}

// NEW creates a MockFSWriteIface.
func (*MockFSWriteIface) NEW(ctrl *gomock.Controller) *MockFSWriteIface {
	return NewMockFSWriteIface(ctrl)
}
