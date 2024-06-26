// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/tq-systems/go-dbus/notification (interfaces: Client)
//
// Generated by this command:
//
//	mockgen --build_flags=--mod=mod -destination=../mocks/notification/mock_notification.go -package=notification github.com/tq-systems/go-dbus/notification Client
//

// Package notification is a generated GoMock package.
package notification

import (
	reflect "reflect"

	notification "github.com/tq-systems/go-dbus/notification"
	gomock "go.uber.org/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// GetPerformance mocks base method.
func (m *MockClient) GetPerformance() (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPerformance")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(uint64)
	ret3, _ := ret[3].(uint64)
	ret4, _ := ret[4].(uint64)
	ret5, _ := ret[5].(uint64)
	ret6, _ := ret[6].(uint64)
	ret7, _ := ret[7].(uint64)
	ret8, _ := ret[8].(error)
	return ret0, ret1, ret2, ret3, ret4, ret5, ret6, ret7, ret8
}

// GetPerformance indicates an expected call of GetPerformance.
func (mr *MockClientMockRecorder) GetPerformance() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPerformance", reflect.TypeOf((*MockClient)(nil).GetPerformance))
}

// Send mocks base method.
func (m *MockClient) Send(arg0 notification.Severity, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockClientMockRecorder) Send(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockClient)(nil).Send), arg0, arg1, arg2)
}

// SendServiceLog mocks base method.
func (m *MockClient) SendServiceLog(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendServiceLog", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendServiceLog indicates an expected call of SendServiceLog.
func (mr *MockClientMockRecorder) SendServiceLog(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendServiceLog", reflect.TypeOf((*MockClient)(nil).SendServiceLog), arg0, arg1)
}
