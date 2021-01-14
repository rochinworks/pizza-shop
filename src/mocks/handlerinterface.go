// Code generated by MockGen. DO NOT EDIT.
// Source: handlerinterface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
)

// MockHandler is a mock of Handler interface
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// BaseHandler mocks base method
func (m *MockHandler) BaseHandler() http.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseHandler")
	ret0, _ := ret[0].(http.HandlerFunc)
	return ret0
}

// BaseHandler indicates an expected call of BaseHandler
func (mr *MockHandlerMockRecorder) BaseHandler() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseHandler", reflect.TypeOf((*MockHandler)(nil).BaseHandler))
}