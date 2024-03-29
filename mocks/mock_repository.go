// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nalcheg/http-checker/repository (interfaces: RepositoryInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	types "github.com/nalcheg/http-checker/types"
	reflect "reflect"
)

// MockRepositoryInterface is a mock of RepositoryInterface interface
type MockRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryInterfaceMockRecorder
}

// MockRepositoryInterfaceMockRecorder is the mock recorder for MockRepositoryInterface
type MockRepositoryInterfaceMockRecorder struct {
	mock *MockRepositoryInterface
}

// NewMockRepositoryInterface creates a new mock instance
func NewMockRepositoryInterface(ctrl *gomock.Controller) *MockRepositoryInterface {
	mock := &MockRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepositoryInterface) EXPECT() *MockRepositoryInterfaceMockRecorder {
	return m.recorder
}

// GetHosts mocks base method
func (m *MockRepositoryInterface) GetHosts() ([]string, error) {
	ret := m.ctrl.Call(m, "GetHosts")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHosts indicates an expected call of GetHosts
func (mr *MockRepositoryInterfaceMockRecorder) GetHosts() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHosts", reflect.TypeOf((*MockRepositoryInterface)(nil).GetHosts))
}

// SaveCheck mocks base method
func (m *MockRepositoryInterface) SaveCheck(arg0 types.Result) error {
	ret := m.ctrl.Call(m, "SaveCheck", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveCheck indicates an expected call of SaveCheck
func (mr *MockRepositoryInterfaceMockRecorder) SaveCheck(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCheck", reflect.TypeOf((*MockRepositoryInterface)(nil).SaveCheck), arg0)
}
