// Code generated by MockGen. DO NOT EDIT.
// Source: common/networkhelper (interfaces: Network)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	net "net"
	reflect "reflect"
)

// MockNetwork is a mock of Network interface
type MockNetwork struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkMockRecorder
}

// MockNetworkMockRecorder is the mock recorder for MockNetwork
type MockNetworkMockRecorder struct {
	mock *MockNetwork
}

// NewMockNetwork creates a new mock instance
func NewMockNetwork(ctrl *gomock.Controller) *MockNetwork {
	mock := &MockNetwork{ctrl: ctrl}
	mock.recorder = &MockNetworkMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetwork) EXPECT() *MockNetworkMockRecorder {
	return m.recorder
}

// AppendSubscriber mocks base method
func (m *MockNetwork) AppendSubscriber() chan net.IP {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendSubscriber")
	ret0, _ := ret[0].(chan net.IP)
	return ret0
}

// AppendSubscriber indicates an expected call of AppendSubscriber
func (mr *MockNetworkMockRecorder) AppendSubscriber() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendSubscriber", reflect.TypeOf((*MockNetwork)(nil).AppendSubscriber))
}

// CheckConnectivity mocks base method
func (m *MockNetwork) CheckConnectivity() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckConnectivity")
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckConnectivity indicates an expected call of CheckConnectivity
func (mr *MockNetworkMockRecorder) CheckConnectivity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckConnectivity", reflect.TypeOf((*MockNetwork)(nil).CheckConnectivity))
}

// GetMACAddress mocks base method
func (m *MockNetwork) GetMACAddress() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMACAddress")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMACAddress indicates an expected call of GetMACAddress
func (mr *MockNetworkMockRecorder) GetMACAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMACAddress", reflect.TypeOf((*MockNetwork)(nil).GetMACAddress))
}

// GetNetInterface mocks base method
func (m *MockNetwork) GetNetInterface() ([]net.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetInterface")
	ret0, _ := ret[0].([]net.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetInterface indicates an expected call of GetNetInterface
func (mr *MockNetworkMockRecorder) GetNetInterface() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetInterface", reflect.TypeOf((*MockNetwork)(nil).GetNetInterface))
}

// GetOutboundIP mocks base method
func (m *MockNetwork) GetOutboundIP() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOutboundIP")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOutboundIP indicates an expected call of GetOutboundIP
func (mr *MockNetworkMockRecorder) GetOutboundIP() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOutboundIP", reflect.TypeOf((*MockNetwork)(nil).GetOutboundIP))
}

// StartNetwork mocks base method
func (m *MockNetwork) StartNetwork() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartNetwork")
}

// StartNetwork indicates an expected call of StartNetwork
func (mr *MockNetworkMockRecorder) StartNetwork() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartNetwork", reflect.TypeOf((*MockNetwork)(nil).StartNetwork))
}
