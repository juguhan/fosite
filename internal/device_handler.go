// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ory/fosite (interfaces: DeviceAuthorizeEndpointHandler,RFC8628UserAuthorizeEndpointHandler)

// Package internal is a generated GoMock package.
package internal

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	fosite "github.com/ory/fosite"
)

// MockDeviceAuthorizeEndpointHandler is a mock of DeviceAuthorizeEndpointHandler interface.
type MockDeviceAuthorizeEndpointHandler struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceAuthorizeEndpointHandlerMockRecorder
}

// MockDeviceAuthorizeEndpointHandlerMockRecorder is the mock recorder for MockDeviceAuthorizeEndpointHandler.
type MockDeviceAuthorizeEndpointHandlerMockRecorder struct {
	mock *MockDeviceAuthorizeEndpointHandler
}

// NewMockDeviceAuthorizeEndpointHandler creates a new mock instance.
func NewMockDeviceAuthorizeEndpointHandler(ctrl *gomock.Controller) *MockDeviceAuthorizeEndpointHandler {
	mock := &MockDeviceAuthorizeEndpointHandler{ctrl: ctrl}
	mock.recorder = &MockDeviceAuthorizeEndpointHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeviceAuthorizeEndpointHandler) EXPECT() *MockDeviceAuthorizeEndpointHandlerMockRecorder {
	return m.recorder
}

// HandleDeviceAuthorizeEndpointRequest mocks base method.
func (m *MockDeviceAuthorizeEndpointHandler) HandleDeviceAuthorizeEndpointRequest(arg0 context.Context, arg1 fosite.DeviceAuthorizeRequester, arg2 fosite.DeviceAuthorizeResponder) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleDeviceAuthorizeEndpointRequest", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleDeviceAuthorizeEndpointRequest indicates an expected call of HandleDeviceAuthorizeEndpointRequest.
func (mr *MockDeviceAuthorizeEndpointHandlerMockRecorder) HandleDeviceAuthorizeEndpointRequest(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleDeviceAuthorizeEndpointRequest", reflect.TypeOf((*MockDeviceAuthorizeEndpointHandler)(nil).HandleDeviceAuthorizeEndpointRequest), arg0, arg1, arg2)
}

// MockRFC8628UserAuthorizeEndpointHandler is a mock of RFC8628UserAuthorizeEndpointHandler interface.
type MockRFC8628UserAuthorizeEndpointHandler struct {
	ctrl     *gomock.Controller
	recorder *MockRFC8628UserAuthorizeEndpointHandlerMockRecorder
}

// MockRFC8628UserAuthorizeEndpointHandlerMockRecorder is the mock recorder for MockRFC8628UserAuthorizeEndpointHandler.
type MockRFC8628UserAuthorizeEndpointHandlerMockRecorder struct {
	mock *MockRFC8628UserAuthorizeEndpointHandler
}

// NewMockRFC8628UserAuthorizeEndpointHandler creates a new mock instance.
func NewMockRFC8628UserAuthorizeEndpointHandler(ctrl *gomock.Controller) *MockRFC8628UserAuthorizeEndpointHandler {
	mock := &MockRFC8628UserAuthorizeEndpointHandler{ctrl: ctrl}
	mock.recorder = &MockRFC8628UserAuthorizeEndpointHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRFC8628UserAuthorizeEndpointHandler) EXPECT() *MockRFC8628UserAuthorizeEndpointHandlerMockRecorder {
	return m.recorder
}

// PopulateRFC8628UserAuthorizeEndpointResponse mocks base method.
func (m *MockRFC8628UserAuthorizeEndpointHandler) PopulateRFC8628UserAuthorizeEndpointResponse(arg0 context.Context, arg1 fosite.DeviceAuthorizeRequester, arg2 fosite.RFC8628UserAuthorizeResponder) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PopulateRFC8628UserAuthorizeEndpointResponse", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PopulateRFC8628UserAuthorizeEndpointResponse indicates an expected call of PopulateRFC8628UserAuthorizeEndpointResponse.
func (mr *MockRFC8628UserAuthorizeEndpointHandlerMockRecorder) PopulateRFC8628UserAuthorizeEndpointResponse(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PopulateRFC8628UserAuthorizeEndpointResponse", reflect.TypeOf((*MockRFC8628UserAuthorizeEndpointHandler)(nil).PopulateRFC8628UserAuthorizeEndpointResponse), arg0, arg1, arg2)
}

// HandleRFC8628UserAuthorizeEndpointRequest mocks base method.
func (m *MockRFC8628UserAuthorizeEndpointHandler) HandleRFC8628UserAuthorizeEndpointRequest(arg0 context.Context, arg1 fosite.DeviceAuthorizeRequester) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleRFC8628UserAuthorizeEndpointRequest", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleRFC8628UserAuthorizeEndpointRequest indicates an expected call of HandleRFC8628UserAuthorizeEndpointRequest.
func (mr *MockRFC8628UserAuthorizeEndpointHandlerMockRecorder) HandleRFC8628UserAuthorizeEndpointRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleRFC8628UserAuthorizeEndpointRequest", reflect.TypeOf((*MockRFC8628UserAuthorizeEndpointHandler)(nil).HandleRFC8628UserAuthorizeEndpointRequest), arg0, arg1)
}
