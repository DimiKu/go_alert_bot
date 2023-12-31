// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package channels is a generated GoMock package.
package channels

import (
	db_actions "go_alert_bot/internal/db_actions"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockChannelRepo is a mock of ChannelRepo interface.
type MockChannelRepo struct {
	ctrl     *gomock.Controller
	recorder *MockChannelRepoMockRecorder
}

// MockChannelRepoMockRecorder is the mock recorder for MockChannelRepo.
type MockChannelRepoMockRecorder struct {
	mock *MockChannelRepo
}

// NewMockChannelRepo creates a new mock instance.
func NewMockChannelRepo(ctrl *gomock.Controller) *MockChannelRepo {
	mock := &MockChannelRepo{ctrl: ctrl}
	mock.recorder = &MockChannelRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChannelRepo) EXPECT() *MockChannelRepoMockRecorder {
	return m.recorder
}

// CreateStdoutChannel mocks base method.
func (m *MockChannelRepo) CreateStdoutChannel(channel db_actions.ChannelDb) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStdoutChannel", channel)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateStdoutChannel indicates an expected call of CreateStdoutChannel.
func (mr *MockChannelRepoMockRecorder) CreateStdoutChannel(channel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStdoutChannel", reflect.TypeOf((*MockChannelRepo)(nil).CreateStdoutChannel), channel)
}

// CreateTelegramChannel mocks base method.
func (m *MockChannelRepo) CreateTelegramChannel(channel db_actions.ChannelDb) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTelegramChannel", channel)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTelegramChannel indicates an expected call of CreateTelegramChannel.
func (mr *MockChannelRepoMockRecorder) CreateTelegramChannel(channel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTelegramChannel", reflect.TypeOf((*MockChannelRepo)(nil).CreateTelegramChannel), channel)
}

// IsExistChannel mocks base method.
func (m *MockChannelRepo) IsExistChannel(channel db_actions.ChannelDb) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsExistChannel", channel)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsExistChannel indicates an expected call of IsExistChannel.
func (mr *MockChannelRepoMockRecorder) IsExistChannel(channel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsExistChannel", reflect.TypeOf((*MockChannelRepo)(nil).IsExistChannel), channel)
}
