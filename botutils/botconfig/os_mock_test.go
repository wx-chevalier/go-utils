// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/wx-chevalier/go-utils/osutils (interfaces: OsClient)

// Package botconfig_test is a generated GoMock package.
package botconfig_test

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOsClient is a mock of OsClient interface
type MockOsClient struct {
	ctrl     *gomock.Controller
	recorder *MockOsClientMockRecorder
}

// MockOsClientMockRecorder is the mock recorder for MockOsClient
type MockOsClientMockRecorder struct {
	mock *MockOsClient
}

// NewMockOsClient creates a new mock instance
func NewMockOsClient(ctrl *gomock.Controller) *MockOsClient {
	mock := &MockOsClient{ctrl: ctrl}
	mock.recorder = &MockOsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOsClient) EXPECT() *MockOsClientMockRecorder {
	return m.recorder
}

// Getenv mocks base method
func (m *MockOsClient) Getenv(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Getenv", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Getenv indicates an expected call of Getenv
func (mr *MockOsClientMockRecorder) Getenv(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Getenv", reflect.TypeOf((*MockOsClient)(nil).Getenv), arg0)
}

// ReadFile mocks base method
func (m *MockOsClient) ReadFile(arg0 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFile", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFile indicates an expected call of ReadFile
func (mr *MockOsClientMockRecorder) ReadFile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFile", reflect.TypeOf((*MockOsClient)(nil).ReadFile), arg0)
}
